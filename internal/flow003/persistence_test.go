package flow003

import (
	"path/filepath"
	"testing"

	"training-review/internal/service"
	"training-review/internal/store"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "persist.db")
	first, err := store.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	clock := service.FixedClock{Value: "2024-01-01T00:00:00Z"}
	registrations := service.NewRegistrationService(first, clock)
	item, err := registrations.Register("go-202", "周七", "zhou@example.test", "平台", []string{"云"})
	if err != nil {
		t.Fatal(err)
	}
	if err := first.Close(); err != nil {
		t.Fatal(err)
	}
	second, err := store.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer second.Close()
	loaded, err := second.GetRegistration(item.ID)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Participant != "周七" || loaded.Status.String() != "draft" {
		t.Fatalf("unexpected reopened value: %+v", loaded)
	}
}

func TestWorkflowImportReport(t *testing.T) {
	db, err := store.Open(filepath.Join(t.TempDir(), "import.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	registrations := service.NewRegistrationService(db, service.FixedClock{Value: "2024-01-01T00:00:00Z"})
	result, err := registrations.Import([]service.ImportRow{{CourseID: "go-1", Participant: "甲", Email: "a@example.test", Department: "研发"}, {CourseID: "", Participant: "乙", Email: "b@example.test"}, {CourseID: "go-2", Participant: "丙", Email: "bad"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Accepted) != 1 || len(result.Rejected) != 2 {
		t.Fatalf("unexpected import counts: %+v", result)
	}
}

func TestSummaryCounts(t *testing.T) {
	db, err := store.Open(filepath.Join(t.TempDir(), "summary.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	registrations := service.NewRegistrationService(db, service.FixedClock{Value: "2024-01-01T00:00:00Z"})
	if _, err = registrations.Register("go-1", "甲", "a@example.test", "研发", nil); err != nil {
		t.Fatal(err)
	}
	if _, err = registrations.Register("go-1", "乙", "b@example.test", "研发", nil); err != nil {
		t.Fatal(err)
	}
	summary, err := registrations.Summary()
	if err != nil || summary.Total != 2 || summary.Courses["go-1"] != 2 {
		t.Fatalf("unexpected summary: %+v %v", summary, err)
	}
}
