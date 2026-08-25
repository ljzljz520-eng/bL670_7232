package service

import (
	"fmt"
	"training-review/internal/domain"
	"training-review/internal/store"
)

type ArchiveService struct {
	store *store.Store
	clock Clock
	ids   *IDGenerator
}

func NewArchiveService(st *store.Store, clock Clock) *ArchiveService {
	return &ArchiveService{store: st, clock: clock, ids: NewIDGenerator("arc")}
}

func (s *ArchiveService) ArchiveRegistration(registrationID, reason string) (domain.Record, error) {
	item, err := s.store.GetRegistration(registrationID)
	if err != nil {
		return domain.Record{}, err
	}
	if item.Status != domain.StatusArchived {
		return domain.Record{}, fmt.Errorf("registration must be archived first")
	}
	record := domain.Record{ID: s.ids.Next(), RegistrationID: registrationID, Kind: "archive", Payload: reason, CreatedAt: s.clock.Now()}
	return record, s.store.SaveRecord(record)
}

func (s *ArchiveService) Attach(registrationID, name, contentType, digest string, size int) (domain.Attachment, error) {
	if size < 0 || name == "" {
		return domain.Attachment{}, fmt.Errorf("attachment metadata is invalid")
	}
	if _, err := s.store.GetRegistration(registrationID); err != nil {
		return domain.Attachment{}, err
	}
	attachment := domain.Attachment{ID: s.ids.Next(), RegistrationID: registrationID, Name: name, ContentType: contentType, Digest: digest, Size: size}
	return attachment, s.store.SaveAttachment(attachment)
}
