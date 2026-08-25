package store

import (
	"errors"
	"fmt"
	"sort"

	bolt "go.etcd.io/bbolt"
	"training-review/internal/domain"
)

var ErrNotFound = errors.New("entity not found")

type Store struct {
	db *bolt.DB
}

func Open(path string) (*Store, error) {
	db, err := bolt.Open(path, 0o600, nil)
	if err != nil {
		return nil, fmt.Errorf("open bbolt database: %w", err)
	}
	s := &Store{db: db}
	if err := db.Update(ensureBuckets); err != nil {
		db.Close()
		return nil, fmt.Errorf("create buckets: %w", err)
	}
	return s, nil
}

func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s *Store) put(bucket, key string, value any) error {
	data, err := encode(value)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(key), data)
	})
}

func (s *Store) get(bucket, key string, target any) error {
	return s.db.View(func(tx *bolt.Tx) error {
		item := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if item == nil {
			return ErrNotFound
		}
		return decode(item, target)
	})
}

func (s *Store) list(bucket string, target func([]byte) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(_, value []byte) error {
			if value == nil {
				return nil
			}
			return target(value)
		})
	})
}

func (s *Store) SaveRegistration(item domain.Registration) error {
	return s.put("registrations", item.ID, item)
}
func (s *Store) GetRegistration(id string) (domain.Registration, error) {
	var item domain.Registration
	err := s.get("registrations", id, &item)
	return item, err
}
func (s *Store) SaveReview(item domain.Review) error     { return s.put("reviews", item.ID, item) }
func (s *Store) SaveRecord(item domain.Record) error     { return s.put("records", item.ID, item) }
func (s *Store) SaveAudit(item domain.AuditEvent) error  { return s.put("audits", item.ID, item) }
func (s *Store) SaveWorkflow(item domain.Workflow) error { return s.put("workflows", item.ID, item) }
func (s *Store) SaveAttachment(item domain.Attachment) error {
	return s.put("attachments", item.ID, item)
}

func (s *Store) ListRegistrations() ([]domain.Registration, error) {
	items := make([]domain.Registration, 0)
	err := s.list("registrations", func(data []byte) error {
		var item domain.Registration
		if err := decode(data, &item); err != nil {
			return err
		}
		items = append(items, item)
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })
	return items, err
}

func (s *Store) ListReviews(registrationID string) ([]domain.Review, error) {
	items := make([]domain.Review, 0)
	err := s.list("reviews", func(data []byte) error {
		var item domain.Review
		if err := decode(data, &item); err != nil {
			return err
		}
		if item.RegistrationID == registrationID {
			items = append(items, item)
		}
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].ReviewedAt < items[j].ReviewedAt })
	return items, err
}

func (s *Store) ListAudits(registrationID string) ([]domain.AuditEvent, error) {
	items := make([]domain.AuditEvent, 0)
	err := s.list("audits", func(data []byte) error {
		var item domain.AuditEvent
		if err := decode(data, &item); err != nil {
			return err
		}
		if registrationID == "" || item.RegistrationID == registrationID {
			items = append(items, item)
		}
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].At < items[j].At })
	return items, err
}
