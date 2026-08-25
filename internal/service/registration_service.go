package service

import (
	"fmt"
	"strings"

	"training-review/internal/domain"
	"training-review/internal/store"
)

type RegistrationService struct {
	store *store.Store
	clock Clock
	ids   *IDGenerator
}

func NewRegistrationService(st *store.Store, clock Clock) *RegistrationService {
	return &RegistrationService{store: st, clock: clock, ids: NewIDGenerator("reg")}
}

func (s *RegistrationService) Register(courseID, participant, email, department string, tags []string) (domain.Registration, error) {
	if strings.TrimSpace(courseID) == "" || strings.TrimSpace(participant) == "" {
		return domain.Registration{}, fmt.Errorf("course and participant are required")
	}
	if !strings.Contains(email, "@") {
		return domain.Registration{}, fmt.Errorf("email is invalid")
	}
	cleanTags := normalizeTags(tags)
	item := domain.Registration{ID: s.ids.Next(), CourseID: courseID, Participant: participant, Email: email, Department: department, SubmittedAt: s.clock.Now(), Status: domain.StatusDraft, Version: 1, Tags: cleanTags}
	if err := ValidateBeforePersist(item); err != nil {
		return domain.Registration{}, err
	}
	if err := s.store.SaveRegistration(item); err != nil {
		return domain.Registration{}, err
	}
	if err := s.record(item.ID, "register", participant); err != nil {
		return domain.Registration{}, err
	}
	return item, nil
}

func normalizeTags(tags []string) []string {
	result := make([]string, 0, len(tags))
	seen := map[string]bool{}
	for _, tag := range tags {
		value := strings.TrimSpace(tag)
		if value != "" && !seen[value] {
			result = append(result, value)
			seen[value] = true
		}
	}
	return result
}

func (s *RegistrationService) Submit(id string) (domain.Registration, error) {
	return s.transition(id, domain.StatusPending, "submit")
}

func (s *RegistrationService) transition(id string, target domain.Status, action string) (domain.Registration, error) {
	item, err := s.store.GetRegistration(id)
	if err != nil {
		return domain.Registration{}, err
	}
	if !domain.CanTransition(item.Status, target) {
		return domain.Registration{}, fmt.Errorf("cannot transition %s to %s", item.Status, target)
	}
	item.Status = target
	item.Version++
	if err := s.store.SaveRegistration(item); err != nil {
		return domain.Registration{}, err
	}
	if err := s.record(id, action, string(target)); err != nil {
		return domain.Registration{}, err
	}
	return item, nil
}

func (s *RegistrationService) Review(id, reviewer string, approve bool, note string) (domain.Registration, domain.Review, error) {
	item, err := s.store.GetRegistration(id)
	if err != nil {
		return domain.Registration{}, domain.Review{}, err
	}
	decision := domain.StatusRejected
	if approve {
		decision = domain.StatusApproved
	}
	if !domain.CanTransition(item.Status, decision) {
		return domain.Registration{}, domain.Review{}, fmt.Errorf("registration is not pending")
	}
	item.Status = decision
	item.Version++
	review := domain.Review{ID: s.ids.Next(), RegistrationID: id, Reviewer: reviewer, Decision: decision, Note: note, ReviewedAt: s.clock.Now()}
	if err := s.store.SaveRegistration(item); err != nil {
		return domain.Registration{}, domain.Review{}, err
	}
	if err := s.store.SaveReview(review); err != nil {
		return domain.Registration{}, domain.Review{}, err
	}
	if err := s.record(id, "review", reviewer); err != nil {
		return domain.Registration{}, domain.Review{}, err
	}
	return item, review, nil
}

func (s *RegistrationService) Archive(id string) (domain.Registration, error) {
	return s.transition(id, domain.StatusArchived, "archive")
}

func (s *RegistrationService) Update(id, courseID, department string, tags []string) (domain.Registration, error) {
	item, err := s.store.GetRegistration(id)
	if err != nil {
		return domain.Registration{}, err
	}
	if domain.IsTerminal(item.Status) {
		return domain.Registration{}, fmt.Errorf("terminal registration cannot be updated")
	}
	if courseID != "" {
		item.CourseID = courseID
	}
	if department != "" {
		item.Department = department
	}
	if tags != nil {
		item.Tags = normalizeTags(tags)
	}
	item.Version++
	if err := s.store.SaveRegistration(item); err != nil {
		return domain.Registration{}, err
	}
	if err := s.record(id, "update", item.Department); err != nil {
		return domain.Registration{}, err
	}
	return item, nil
}

func (s *RegistrationService) Search(filter domain.SearchFilter) ([]domain.Registration, error) {
	return s.store.Search(filter)
}

func (s *RegistrationService) Get(id string) (domain.Registration, error) {
	return s.store.GetRegistration(id)
}

func (s *RegistrationService) Summary() (domain.EnrollmentSummary, error) { return s.store.Summary() }

func (s *RegistrationService) History(id string) ([]domain.AuditEvent, error) {
	return s.store.AuditTrail(id)
}

func (s *RegistrationService) Reviews(id string) ([]domain.Review, error) {
	return s.store.ReviewsFor(id)
}

func (s *RegistrationService) Attachments(id string) ([]domain.Attachment, error) {
	return s.store.AttachmentsFor(id)
}

func (s *RegistrationService) Validate(id string) ([]domain.ValidationIssue, error) {
	item, err := s.store.GetRegistration(id)
	if err != nil {
		return nil, err
	}
	return domain.ValidateRegistration(item), nil
}

func (s *RegistrationService) CanEdit(id string) (bool, error) {
	item, err := s.store.GetRegistration(id)
	if err != nil {
		return false, err
	}
	return !domain.IsTerminal(item.Status), nil
}

func (s *RegistrationService) record(registrationID, action, detail string) error {
	event := domain.AuditEvent{ID: s.ids.Next(), RegistrationID: registrationID, Action: action, Actor: "system", Detail: detail, At: s.clock.Now()}
	return s.store.SaveAudit(event)
}
