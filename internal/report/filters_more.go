package report

import "training-review/internal/domain"

func ByDepartment(items []domain.Registration, department string) []domain.Registration {
	result := make([]domain.Registration, 0)
	for _, item := range items {
		if item.Department == department {
			result = append(result, item)
		}
	}
	return result
}

func ByCourse(items []domain.Registration, courseID string) []domain.Registration {
	result := make([]domain.Registration, 0)
	for _, item := range items {
		if item.CourseID == courseID {
			result = append(result, item)
		}
	}
	return result
}

func ByVersionAtLeast(items []domain.Registration, version int) []domain.Registration {
	result := make([]domain.Registration, 0)
	for _, item := range items {
		if item.Version >= version {
			result = append(result, item)
		}
	}
	return result
}

func WithEmailDomain(items []domain.Registration, domainName string) []domain.Registration {
	result := make([]domain.Registration, 0)
	suffix := "@" + domainName
	for _, item := range items {
		if len(item.Email) >= len(suffix) && item.Email[len(item.Email)-len(suffix):] == suffix {
			result = append(result, item)
		}
	}
	return result
}
