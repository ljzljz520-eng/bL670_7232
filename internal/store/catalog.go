package store

import (
	"sort"
	"training-review/internal/domain"
)

func (s *Store) SaveCourse(course domain.Course) error { return s.put("courses", course.ID, course) }

func (s *Store) GetCourse(id string) (domain.Course, error) {
	var item domain.Course
	err := s.get("courses", id, &item)
	return item, err
}

func (s *Store) ListCourses() ([]domain.Course, error) {
	items := make([]domain.Course, 0)
	err := s.list("courses", func(data []byte) error {
		var item domain.Course
		if err := decode(data, &item); err != nil {
			return err
		}
		items = append(items, item)
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })
	return items, err
}

func (s *Store) SaveReviewer(reviewer domain.Reviewer) error {
	return s.put("reviewers", reviewer.ID, reviewer)
}

func (s *Store) ListReviewers() ([]domain.Reviewer, error) {
	items := make([]domain.Reviewer, 0)
	err := s.list("reviewers", func(data []byte) error {
		var item domain.Reviewer
		if err := decode(data, &item); err != nil {
			return err
		}
		items = append(items, item)
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })
	return items, err
}

func (s *Store) SaveNotification(notification domain.Notification) error {
	return s.put("notifications", notification.ID, notification)
}

func (s *Store) ListNotifications(registrationID string) ([]domain.Notification, error) {
	items := make([]domain.Notification, 0)
	err := s.list("notifications", func(data []byte) error {
		var item domain.Notification
		if err := decode(data, &item); err != nil {
			return err
		}
		if registrationID == "" || item.RegistrationID == registrationID {
			items = append(items, item)
		}
		return nil
	})
	sort.Slice(items, func(i, j int) bool { return items[i].SentAt < items[j].SentAt })
	return items, err
}
