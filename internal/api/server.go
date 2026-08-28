package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"training-review/internal/domain"
	"training-review/internal/service"
)

type Server struct {
	registrations *service.RegistrationService
	workflows     *service.WorkflowService
	archive       *service.ArchiveService
}

func NewServer(registrations *service.RegistrationService, workflows *service.WorkflowService, archive *service.ArchiveService) *Server {
	return &Server{registrations: registrations, workflows: workflows, archive: archive}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.health)
	mux.HandleFunc("/registrations", s.registrationsEndpoint)
	mux.HandleFunc("/registrations/", s.registrationEndpoint)
	mux.HandleFunc("/summary", s.summary)
	return requestLogger(mux)
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-Path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

type registrationRequest struct {
	CourseID    string   `json:"course_id"`
	Participant string   `json:"participant"`
	Email       string   `json:"email"`
	Department  string   `json:"department"`
	Tags        []string `json:"tags"`
}

func (s *Server) registrationsEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.createRegistration(w, r)
		return
	}
	if r.Method == http.MethodGet {
		s.searchRegistrations(w, r)
		return
	}
	writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

func (s *Server) createRegistration(w http.ResponseWriter, r *http.Request) {
	var request registrationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	item, err := s.registrations.Register(request.CourseID, request.Participant, request.Email, request.Department, request.Tags)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func (s *Server) searchRegistrations(w http.ResponseWriter, r *http.Request) {
	filter := domain.SearchFilter{CourseID: r.URL.Query().Get("course_id"), Participant: r.URL.Query().Get("participant"), Department: r.URL.Query().Get("department"), Tag: r.URL.Query().Get("tag")}
	filter.Status = domain.NormalizeStatus(r.URL.Query().Get("status"))
	if r.URL.Query().Get("status") == "" {
		filter.Status = ""
	}
	items, err := s.registrations.Search(filter)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (s *Server) registrationEndpoint(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/registrations/")
	if id == "" {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "missing id"})
		return
	}
	if r.Method == http.MethodGet {
		item, err := s.registrations.Get(id)
		if err != nil {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, item)
		return
	}
	if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/submit") {
		s.submit(w, r, strings.TrimSuffix(id, "/submit"))
		return
	}
	writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

func (s *Server) submit(w http.ResponseWriter, _ *http.Request, id string) {
	item, err := s.registrations.Submit(id)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (s *Server) summary(w http.ResponseWriter, _ *http.Request) {
	summary, err := s.registrations.Summary()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, summary)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
