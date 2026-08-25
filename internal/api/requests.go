package api

import (
	"encoding/json"
	"net/http"
	"training-review/internal/service"
)

type reviewRequest struct {
	Reviewer string `json:"reviewer"`
	Approve  bool   `json:"approve"`
	Note     string `json:"note"`
}
type updateRequest struct {
	CourseID   string   `json:"course_id"`
	Department string   `json:"department"`
	Tags       []string `json:"tags"`
}

func decodeBody(r *http.Request, target any) error { return json.NewDecoder(r.Body).Decode(target) }

func applyUpdate(svc *service.RegistrationService, id string, request updateRequest) (any, error) {
	return svc.Update(id, request.CourseID, request.Department, request.Tags)
}

func validateReview(request reviewRequest) bool { return request.Reviewer != "" }
