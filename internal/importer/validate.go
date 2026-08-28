package importer

import (
	"fmt"
	"strings"
	"training-review/internal/service"
)

type RowIssue struct {
	Line    int
	Field   string
	Message string
}

func (i RowIssue) String() string { return fmt.Sprintf("line %d %s: %s", i.Line, i.Field, i.Message) }

func ValidateRows(rows []service.ImportRow) []RowIssue {
	issues := make([]RowIssue, 0)
	for index, row := range rows {
		line := index + 2
		if strings.TrimSpace(row.CourseID) == "" {
			issues = append(issues, RowIssue{line, "course_id", "required"})
		}
		if strings.TrimSpace(row.Participant) == "" {
			issues = append(issues, RowIssue{line, "participant", "required"})
		}
		if !strings.Contains(row.Email, "@") {
			issues = append(issues, RowIssue{line, "email", "invalid"})
		}
		if strings.TrimSpace(row.Department) == "" {
			issues = append(issues, RowIssue{line, "department", "required"})
		}
	}
	return issues
}

func SummarizeIssues(issues []RowIssue) map[string]int {
	result := make(map[string]int)
	for _, issue := range issues {
		result[issue.Field]++
	}
	return result
}

func IssueText(issues []RowIssue) string {
	values := make([]string, 0, len(issues))
	for _, issue := range issues {
		values = append(values, issue.String())
	}
	return strings.Join(values, "; ")
}

func IsValidRow(row service.ImportRow) bool { return len(ValidateRows([]service.ImportRow{row})) == 0 }
