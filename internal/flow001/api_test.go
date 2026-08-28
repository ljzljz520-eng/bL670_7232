package flow001

import (
	"net/http/httptest"
	"strings"
	"testing"
	"training-review/internal/api"
	"training-review/internal/service"
	"training-review/internal/store"
)

func TestHealthEndpoint(t *testing.T) {
	db, err := store.Open(t.TempDir() + "/api.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	clock := service.FixedClock{Value: "2024-01-01T00:00:00Z"}
	server := api.NewServer(service.NewRegistrationService(db, clock), service.NewWorkflowService(db, clock), service.NewArchiveService(db, clock))
	record := httptest.NewRecorder()
	server.Handler().ServeHTTP(record, httptest.NewRequest("GET", "/health", nil))
	if record.Code != 200 || !strings.Contains(record.Body.String(), "ok") {
		t.Fatalf("unexpected health response: %d %s", record.Code, record.Body.String())
	}
}
