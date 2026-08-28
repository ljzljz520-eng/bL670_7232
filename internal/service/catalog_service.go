package service

import (
	"fmt"
	"strings"
	"training-review/internal/domain"
	"training-review/internal/store"
)

type CatalogService struct {
	store *store.Store
	ids   *IDGenerator
}

func NewCatalogService(st *store.Store) *CatalogService {
	return &CatalogService{store: st, ids: NewIDGenerator("course")}
}

func (s *CatalogService) CreateCourse(title, description, department string, capacity int, tags []string) (domain.Course, error) {
	if strings.TrimSpace(title) == "" || strings.TrimSpace(department) == "" {
		return domain.Course{}, fmt.Errorf("title and department are required")
	}
	if capacity < 1 {
		return domain.Course{}, fmt.Errorf("capacity must be positive")
	}
	course := domain.Course{ID: s.ids.Next(), Title: title, Description: description, Department: department, Capacity: capacity, Active: true, Tags: normalizeTags(tags)}
	return course, s.store.SaveCourse(course)
}

func (s *CatalogService) DeactivateCourse(id string) (domain.Course, error) {
	course, err := s.store.GetCourse(id)
	if err != nil {
		return domain.Course{}, err
	}
	if !course.Active {
		return course, nil
	}
	course.Active = false
	return course, s.store.SaveCourse(course)
}

func (s *CatalogService) ActivateCourse(id string) (domain.Course, error) {
	course, err := s.store.GetCourse(id)
	if err != nil {
		return domain.Course{}, err
	}
	if course.Active {
		return course, nil
	}
	course.Active = true
	return course, s.store.SaveCourse(course)
}

func (s *CatalogService) FindByTag(tag string) ([]domain.Course, error) {
	courses, err := s.store.ListCourses()
	if err != nil {
		return nil, err
	}
	result := make([]domain.Course, 0)
	for _, course := range courses {
		if course.MatchesTag(tag) && course.Active {
			result = append(result, course)
		}
	}
	return result, nil
}

func (s *CatalogService) Available(courseID string, current int) (bool, error) {
	course, err := s.store.GetCourse(courseID)
	if err != nil {
		return false, err
	}
	return course.CanAccept(current), nil
}
