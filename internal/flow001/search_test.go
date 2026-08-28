package flow001

import (
	"path/filepath"
	"testing"
	"training-review/internal/domain"
	"training-review/internal/service"
	"training-review/internal/store"
)

func TestWorkflowSearchUpdatePublish(t *testing.T) {
	db, err := store.Open(filepath.Join(t.TempDir(), "search.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	svc := service.NewRegistrationService(db, service.FixedClock{Value: "2024-01-01T00:00:00Z"})
	first, err := svc.Register("go-1", "甲", "a@example.test", "研发", []string{"backend"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err = svc.Register("go-2", "乙", "b@example.test", "销售", []string{"sales"}); err != nil {
		t.Fatal(err)
	}
	items, err := svc.Search(domain.SearchFilter{Tag: "backend"})
	if err != nil || len(items) != 1 || items[0].ID != first.ID {
		t.Fatalf("search mismatch: %+v %v", items, err)
	}
	updated, err := svc.Update(first.ID, "", "产品", []string{"backend", "leadership"})
	if err != nil {
		t.Fatal(err)
	}
	if updated.Version != 2 || updated.Department != "产品" {
		t.Fatalf("update mismatch: %+v", updated)
	}
}
