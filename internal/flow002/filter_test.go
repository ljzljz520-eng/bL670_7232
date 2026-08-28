package flow002

import "testing"

func TestHandlerDefaultValues(t *testing.T) {
	handler := NewHandler()
	item := handler.Prepare("", "")
	if item.Number != "unassigned" || item.Participant != "unknown" {
		t.Fatal("defaults were not applied")
	}
}
