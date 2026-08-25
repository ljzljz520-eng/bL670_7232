package flow001

import (
	"path/filepath"
	"testing"

	"training-review/internal/domain"
	"training-review/internal/service"
	"training-review/internal/store"
)

func openServices(t *testing.T) (*service.RegistrationService, *service.ArchiveService, *store.Store) {
	t.Helper()
	db, err := store.Open(filepath.Join(t.TempDir(), "workflow.db"))
	if err != nil {
		t.Fatal(err)
	}
	clock := service.FixedClock{Value: "2024-01-01T00:00:00Z"}
	return service.NewRegistrationService(db, clock), service.NewArchiveService(db, clock), db
}

func TestWorkflowCreateReviewArchive(t *testing.T) {
	registrations, archives, db := openServices(t)
	defer db.Close()
	item, err := registrations.Register("go-101", "林梅", "lin@example.test", "研发", []string{"核心", "后端"})
	if err != nil {
		t.Fatal(err)
	}
	if item.Status != domain.StatusDraft {
		t.Fatalf("unexpected status %s", item.Status)
	}
	if _, err = registrations.Submit(item.ID); err != nil {
		t.Fatal(err)
	}
	approved, review, err := registrations.Review(item.ID, "审核员A", true, "资料完整")
	if err != nil {
		t.Fatal(err)
	}
	if approved.Status != domain.StatusApproved || review.Decision != domain.StatusApproved {
		t.Fatalf("review did not approve")
	}
	if _, err = registrations.Archive(item.ID); err != nil {
		t.Fatal(err)
	}
	record, err := archives.ArchiveRegistration(item.ID, "季度归档")
	if err != nil {
		t.Fatal(err)
	}
	if record.RegistrationID != item.ID {
		t.Fatalf("record linked to wrong registration")
	}
}

func TestRegistrationValidation(t *testing.T) {
	registrations, _, db := openServices(t)
	defer db.Close()
	if _, err := registrations.Register("", "", "bad", "", nil); err == nil {
		t.Fatal("expected validation error")
	}
	if _, err := registrations.Register("go-101", "林梅", "bad", "研发", nil); err == nil {
		t.Fatal("expected email error")
	}
}

func TestReviewRejectsIllegalTransition(t *testing.T) {
	registrations, _, db := openServices(t)
	defer db.Close()
	item, err := registrations.Register("go-101", "林梅", "lin@example.test", "研发", nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err = registrations.Review(item.ID, "审核员A", true, "未提交"); err == nil {
		t.Fatal("expected transition error")
	}
}
