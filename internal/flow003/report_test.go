package flow003

import (
	"strings"
	"testing"
	"training-review/internal/domain"
	"training-review/internal/report"
)

func TestReportCSV(t *testing.T) {
	text := report.CSV([]domain.Registration{{ID: "r-1", CourseID: "go", Participant: "甲", Email: "a@example.test", Status: domain.StatusDraft, Version: 1, Tags: []string{"x"}}})
	if !strings.Contains(text, "id,course") || !strings.Contains(text, "r-1,go") {
		t.Fatalf("unexpected csv: %s", text)
	}
}
