package report

import (
	"sort"
	"training-review/internal/domain"
)

type DepartmentMetric struct {
	Department string
	Total      int
	Pending    int
	Approved   int
	Rejected   int
}

func DepartmentMetrics(items []domain.Registration) []DepartmentMetric {
	metrics := make(map[string]DepartmentMetric)
	for _, item := range items {
		metric := metrics[item.Department]
		metric.Department = item.Department
		metric.Total++
		switch item.Status {
		case domain.StatusPending:
			metric.Pending++
		case domain.StatusApproved:
			metric.Approved++
		case domain.StatusRejected:
			metric.Rejected++
		}
		metrics[item.Department] = metric
	}
	result := make([]DepartmentMetric, 0, len(metrics))
	for _, metric := range metrics {
		result = append(result, metric)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Department < result[j].Department })
	return result
}

func CompletionRate(items []domain.Registration) float64 {
	if len(items) == 0 {
		return 0
	}
	completed := 0
	for _, item := range items {
		if item.Status == domain.StatusApproved || item.Status == domain.StatusArchived {
			completed++
		}
	}
	return float64(completed) / float64(len(items))
}

func StatusCounts(items []domain.Registration) map[domain.Status]int {
	counts := make(map[domain.Status]int)
	for _, item := range items {
		counts[item.Status]++
	}
	return counts
}

func CourseCounts(items []domain.Registration) map[string]int {
	counts := make(map[string]int)
	for _, item := range items {
		counts[item.CourseID]++
	}
	return counts
}

func UniqueDepartments(items []domain.Registration) []string {
	seen := map[string]bool{}
	for _, item := range items {
		seen[item.Department] = true
	}
	result := make([]string, 0, len(seen))
	for department := range seen {
		result = append(result, department)
	}
	sort.Strings(result)
	return result
}
