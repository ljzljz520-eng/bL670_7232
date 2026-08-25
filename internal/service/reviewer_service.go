package service

import (
	"fmt"
	"strings"
	"training-review/internal/domain"
	"training-review/internal/store"
)

type ReviewerService struct {
	store *store.Store
	ids   *IDGenerator
}

func NewReviewerService(st *store.Store) *ReviewerService {
	return &ReviewerService{store: st, ids: NewIDGenerator("reviewer")}
}

func (s *ReviewerService) Register(name, department string, specialties []string) (domain.Reviewer, error) {
	if strings.TrimSpace(name) == "" {
		return domain.Reviewer{}, fmt.Errorf("reviewer name is required")
	}
	item := domain.Reviewer{ID: s.ids.Next(), Name: name, Department: department, Specialties: normalizeTags(specialties), Enabled: true}
	return item, s.store.SaveReviewer(item)
}

func (s *ReviewerService) Disable(id string) (domain.Reviewer, error) {
	items, err := s.store.ListReviewers()
	if err != nil {
		return domain.Reviewer{}, err
	}
	for _, item := range items {
		if item.ID == id {
			item.Enabled = false
			return item, s.store.SaveReviewer(item)
		}
	}
	return domain.Reviewer{}, store.ErrNotFound
}

func (s *ReviewerService) Assign(department string) (domain.Reviewer, error) {
	items, err := s.store.ListReviewers()
	if err != nil {
		return domain.Reviewer{}, err
	}
	for _, item := range items {
		if item.CanReview(department) {
			return item, nil
		}
	}
	return domain.Reviewer{}, fmt.Errorf("no reviewer available for %s", department)
}
