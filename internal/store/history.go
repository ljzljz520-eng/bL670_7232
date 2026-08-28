package store

import (
	"sort"
	"training-review/internal/domain"
)

func (s *Store) ReviewsFor(registrationID string) ([]domain.Review, error) {
	return s.ListReviews(registrationID)
}

func (s *Store) AuditTrail(registrationID string) ([]domain.AuditEvent, error) {
	return s.ListAudits(registrationID)
}

func (s *Store) LatestAudit(registrationID string) (domain.AuditEvent, error) {
	items, err := s.ListAudits(registrationID)
	if err != nil {
		return domain.AuditEvent{}, err
	}
	if len(items) == 0 {
		return domain.AuditEvent{}, ErrNotFound
	}
	sort.Slice(items, func(i, j int) bool { return items[i].At > items[j].At })
	return items[0], nil
}

func (s *Store) WorkflowsFor(registrationID string) ([]domain.Workflow, error) {
	items := make([]domain.Workflow, 0)
	err := s.list("workflows", func(data []byte) error {
		var item domain.Workflow
		if err := decode(data, &item); err != nil {
			return err
		}
		if item.RegistrationID == registrationID {
			items = append(items, item)
		}
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].UpdatedAt < items[j].UpdatedAt })
	return items, err
}

func (s *Store) RecordsFor(registrationID string) ([]domain.Record, error) {
	items := make([]domain.Record, 0)
	err := s.list("records", func(data []byte) error {
		var item domain.Record
		if err := decode(data, &item); err != nil {
			return err
		}
		if item.RegistrationID == registrationID {
			items = append(items, item)
		}
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].CreatedAt < items[j].CreatedAt })
	return items, err
}

func (s *Store) AttachmentsFor(registrationID string) ([]domain.Attachment, error) {
	items := make([]domain.Attachment, 0)
	err := s.list("attachments", func(data []byte) error {
		var item domain.Attachment
		if err := decode(data, &item); err != nil {
			return err
		}
		if item.RegistrationID == registrationID {
			items = append(items, item)
		}
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })
	return items, err
}
