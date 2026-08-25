package service

import (
	"sort"
	"training-review/internal/domain"
)

type ServiceMetrics struct {
	Total       int
	Pending     int
	Approved    int
	Rejected    int
	Archived    int
	Departments []string
}

func BuildMetrics(items []domain.Registration) ServiceMetrics {
	metrics := ServiceMetrics{Total: len(items)}
	departments := map[string]bool{}
	for _, item := range items {
		departments[item.Department] = true
		switch item.Status {
		case domain.StatusPending:
			metrics.Pending++
		case domain.StatusApproved:
			metrics.Approved++
		case domain.StatusRejected:
			metrics.Rejected++
		case domain.StatusArchived:
			metrics.Archived++
		}
	}
	metrics.Departments = make([]string, 0, len(departments))
	for department := range departments {
		metrics.Departments = append(metrics.Departments, department)
	}
	sort.Strings(metrics.Departments)
	return metrics
}

func (m ServiceMetrics) Completed() int { return m.Approved + m.Archived }

func (m ServiceMetrics) Open() int { return m.Total - m.Completed() - m.Rejected }

func (m ServiceMetrics) Ratio() float64 {
	if m.Total == 0 {
		return 0
	}
	return float64(m.Completed()) / float64(m.Total)
}

func (m ServiceMetrics) HasPending() bool { return m.Pending > 0 }

func (m ServiceMetrics) DepartmentCount() int { return len(m.Departments) }

func (s *RegistrationService) Metrics() (ServiceMetrics, error) {
	items, err := s.store.ListRegistrations()
	if err != nil {
		return ServiceMetrics{}, err
	}
	return BuildMetrics(items), nil
}
