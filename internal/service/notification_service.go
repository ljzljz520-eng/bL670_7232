package service

import (
	"fmt"
	"strings"
	"training-review/internal/domain"
	"training-review/internal/store"
)

type NotificationService struct {
	store *store.Store
	clock Clock
	ids   *IDGenerator
}

func NewNotificationService(st *store.Store, clock Clock) *NotificationService {
	return &NotificationService{store: st, clock: clock, ids: NewIDGenerator("notice")}
}

func (s *NotificationService) Queue(registration domain.Registration, subject, body string) (domain.Notification, error) {
	if registration.ID == "" || !strings.Contains(registration.Email, "@") {
		return domain.Notification{}, fmt.Errorf("registration recipient is invalid")
	}
	if strings.TrimSpace(subject) == "" || strings.TrimSpace(body) == "" {
		return domain.Notification{}, fmt.Errorf("notification content is required")
	}
	notification := domain.Notification{ID: s.ids.Next(), RegistrationID: registration.ID, Channel: "email", Recipient: registration.Email, Subject: subject, Body: body, SentAt: s.clock.Now()}
	return notification, s.store.SaveNotification(notification)
}

func (s *NotificationService) QueueStatus(registration domain.Registration) (domain.Notification, error) {
	return s.Queue(registration, "培训报名状态更新", "当前状态："+domain.StatusLabel(registration.Status))
}

func (s *NotificationService) Pending(registrationID string) ([]domain.Notification, error) {
	return s.store.ListNotifications(registrationID)
}
