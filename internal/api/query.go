package api

import (
	"net/http"
	"training-review/internal/domain"
)

func queryFilter(r *http.Request) domain.SearchFilter {
	filter := domain.SearchFilter{CourseID: r.URL.Query().Get("course_id"), Participant: r.URL.Query().Get("participant"), Department: r.URL.Query().Get("department"), Tag: r.URL.Query().Get("tag")}
	if value := r.URL.Query().Get("status"); value != "" {
		filter.Status = domain.NormalizeStatus(value)
	}
	return filter
}

func wantsCSV(r *http.Request) bool {
	return r.URL.Query().Get("format") == "csv" || r.Header.Get("Accept") == "text/csv"
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

func writeMethodError(w http.ResponseWriter) {
	writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

func readID(path string) string {
	for len(path) > 0 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
