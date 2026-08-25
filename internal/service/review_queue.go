package service

import (
	"fmt"
	"training-review/internal/domain"
)

type ReviewQueue struct {
	registrations *RegistrationService
	reviewers     *ReviewerService
}

func NewReviewQueue(registrations *RegistrationService, reviewers *ReviewerService) *ReviewQueue {
	return &ReviewQueue{registrations: registrations, reviewers: reviewers}
}

func (q *ReviewQueue) Next(department string) (domain.Registration, domain.Reviewer, error) {
	items, err := q.registrations.Search(domain.SearchFilter{Department: department, Status: domain.StatusPending})
	if err != nil {
		return domain.Registration{}, domain.Reviewer{}, err
	}
	if len(items) == 0 {
		return domain.Registration{}, domain.Reviewer{}, fmt.Errorf("review queue is empty")
	}
	reviewer, err := q.reviewers.Assign(department)
	if err != nil {
		return domain.Registration{}, domain.Reviewer{}, err
	}
	return items[0], reviewer, nil
}

func (q *ReviewQueue) Size(department string) (int, error) {
	items, err := q.registrations.Search(domain.SearchFilter{Department: department, Status: domain.StatusPending})
	return len(items), err
}
