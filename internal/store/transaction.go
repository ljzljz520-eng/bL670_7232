package store

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"training-review/internal/domain"
)

func (s *Store) SaveRegistrationWithAudit(item domain.Registration, audit domain.AuditEvent) error {
	registrationData, err := encode(item)
	if err != nil {
		return err
	}
	auditData, err := encode(audit)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket([]byte("registrations")).Put([]byte(item.ID), registrationData); err != nil {
			return fmt.Errorf("save registration: %w", err)
		}
		return tx.Bucket([]byte("audits")).Put([]byte(audit.ID), auditData)
	})
}

func (s *Store) Count(bucket string) (int, error) {
	count := 0
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return ErrNotFound
		}
		return b.ForEach(func(key, value []byte) error {
			if key != nil && value != nil {
				count++
			}
			return nil
		})
	})
	return count, err
}

func (s *Store) Counts() (map[string]int, error) {
	counts := make(map[string]int, len(bucketNames))
	for _, name := range bucketNames {
		count, err := s.Count(string(name))
		if err != nil {
			return nil, err
		}
		counts[string(name)] = count
	}
	return counts, nil
}
