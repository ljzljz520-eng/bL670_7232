package report

import (
	"sort"
	"training-review/internal/domain"
)

func SortByStatus(items []domain.Registration) []domain.Registration {
	result := append([]domain.Registration(nil), items...)
	sort.SliceStable(result, func(i, j int) bool {
		if result[i].Status == result[j].Status {
			return result[i].ID < result[j].ID
		}
		return result[i].Status < result[j].Status
	})
	return result
}

func CountTags(items []domain.Registration) map[string]int {
	counts := map[string]int{}
	for _, item := range items {
		for _, tag := range item.Tags {
			counts[tag]++
		}
	}
	return counts
}

func Pending(items []domain.Registration) []domain.Registration {
	result := make([]domain.Registration, 0)
	for _, item := range items {
		if item.Status == domain.StatusPending {
			result = append(result, item)
		}
	}
	return result
}
