package importer

import (
	"sort"
	"training-review/internal/domain"
)

type ImportPlan struct {
	Rows        int
	Valid       int
	Invalid     int
	Departments []string
}

func BuildPlan(items []domain.Registration) ImportPlan {
	departments := make(map[string]bool)
	valid := 0
	for _, item := range items {
		departments[item.Department] = true
		if len(domain.ValidateRegistration(item)) == 0 {
			valid++
		}
	}
	result := ImportPlan{Rows: len(items), Valid: valid, Invalid: len(items) - valid, Departments: make([]string, 0, len(departments))}
	for department := range departments {
		result.Departments = append(result.Departments, department)
	}
	sort.Strings(result.Departments)
	return result
}

func (p ImportPlan) Complete() bool { return p.Rows > 0 && p.Invalid == 0 }

func (p ImportPlan) CompletionMessage() string {
	if p.Complete() {
		return "导入计划可执行"
	}
	return "导入计划需要修正"
}
