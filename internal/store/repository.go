package store

import (
	"sort"
	"training-review/internal/domain"
)

func (s *Store) ReplaceTags(id string, tags []string) (domain.Registration, error) {
	item, err := s.GetRegistration(id)
	if err != nil {
		return domain.Registration{}, err
	}
	item.Tags = append([]string(nil), tags...)
	item.Version++
	if err := s.SaveRegistration(item); err != nil {
		return domain.Registration{}, err
	}
	return item, nil
}

func (s *Store) RegistrationsByVersion(minimum int) ([]domain.Registration, error) {
	items, err := s.ListRegistrations()
	if err != nil {
		return nil, err
	}
	result := make([]domain.Registration, 0)
	for _, item := range items {
		if item.Version >= minimum {
			result = append(result, item)
		}
	}
	return result, nil
}

func (s *Store) RegistrationsByStatus(status domain.Status) ([]domain.Registration, error) {
	items, err := s.ListRegistrations()
	if err != nil {
		return nil, err
	}
	result := make([]domain.Registration, 0)
	for _, item := range items {
		if item.Status == status {
			result = append(result, item)
		}
	}
	return result, nil
}

func (s *Store) LatestRegistration() (domain.Registration, error) {
	items, err := s.ListRegistrations()
	if err != nil {
		return domain.Registration{}, err
	}
	if len(items) == 0 {
		return domain.Registration{}, ErrNotFound
	}
	sort.Slice(items, func(i, j int) bool { return items[i].SubmittedAt > items[j].SubmittedAt })
	return items[0], nil
}

func (s *Store) SaveAll(items []domain.Registration) error {
	for _, item := range items {
		if err := s.SaveRegistration(item); err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) SaveReviews(items []domain.Review) error {
	for _, item := range items {
		if err := s.SaveReview(item); err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) SaveAudits(items []domain.AuditEvent) error {
	for _, item := range items {
		if err := s.SaveAudit(item); err != nil {
			return err
		}
	}
	return nil
}
