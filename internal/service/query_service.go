package service

import (
	"sort"
	"training-review/internal/domain"
)

type QueryService struct{ registrations *RegistrationService }

func NewQueryService(registrations *RegistrationService) *QueryService {
	return &QueryService{registrations: registrations}
}

func (s *QueryService) PendingByDepartment(department string) ([]domain.Registration, error) {
	items, err := s.registrations.Search(domain.SearchFilter{Department: department, Status: domain.StatusPending})
	if err != nil {
		return nil, err
	}
	sort.Slice(items, func(i, j int) bool { return items[i].SubmittedAt < items[j].SubmittedAt })
	return items, nil
}

func (s *QueryService) ParticipantHistory(participant string) ([]domain.Registration, error) {
	items, err := s.registrations.Search(domain.SearchFilter{Participant: participant})
	if err != nil {
		return nil, err
	}
	sort.Slice(items, func(i, j int) bool { return items[i].Version < items[j].Version })
	return items, nil
}

func (s *QueryService) CourseRoster(courseID string) ([]domain.Registration, error) {
	return s.registrations.Search(domain.SearchFilter{CourseID: courseID})
}

func (s *QueryService) HasPending(courseID string) (bool, error) {
	items, err := s.registrations.Search(domain.SearchFilter{CourseID: courseID, Status: domain.StatusPending})
	return len(items) > 0, err
}
