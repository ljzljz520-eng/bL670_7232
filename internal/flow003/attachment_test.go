package flow003

import (
	"path/filepath"
	"testing"
	"training-review/internal/service"
	"training-review/internal/store"
)

func TestAttachmentLifecycle(t *testing.T) {
	db, err := store.Open(filepath.Join(t.TempDir(), "attachment.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	clock := service.FixedClock{Value: "2024-01-01T00:00:00Z"}
	registrations := service.NewRegistrationService(db, clock)
	item, err := registrations.Register("go", "甲", "a@example.test", "研发", nil)
	if err != nil {
		t.Fatal(err)
	}
	archive := service.NewArchiveService(db, clock)
	attachment, err := archive.Attach(item.ID, "proof.pdf", "application/pdf", "digest", 12)
	if err != nil || attachment.Size != 12 {
		t.Fatalf("attachment failed: %+v %v", attachment, err)
	}
}
