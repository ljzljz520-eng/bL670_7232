package store

import "training-review/internal/domain"

func matches(item domain.Registration, filter domain.SearchFilter) bool {
	if filter.CourseID != "" && item.CourseID != filter.CourseID {
		return false
	}
	if filter.Participant != "" && item.Participant != filter.Participant {
		return false
	}
	if filter.Department != "" && item.Department != filter.Department {
		return false
	}
	if filter.Status != "" && item.Status != filter.Status {
		return false
	}
	if filter.Tag != "" {
		found := false
		for _, tag := range item.Tags {
			if tag == filter.Tag {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (s *Store) Search(filter domain.SearchFilter) ([]domain.Registration, error) {
	all, err := s.ListRegistrations()
	if err != nil {
		return nil, err
	}
	result := make([]domain.Registration, 0, len(all))
	for _, item := range all {
		if matches(item, filter) {
			result = append(result, item)
		}
	}
	return result, nil
}

func (s *Store) Summary() (domain.EnrollmentSummary, error) {
	items, err := s.ListRegistrations()
	if err != nil {
		return domain.EnrollmentSummary{}, err
	}
	summary := domain.EnrollmentSummary{ByStatus: map[domain.Status]int{}, Courses: map[string]int{}}
	for _, item := range items {
		summary.Total++
		summary.ByStatus[item.Status]++
		summary.Courses[item.CourseID]++
	}
	return summary, nil
}
