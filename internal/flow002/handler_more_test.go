package flow002

import "testing"

func TestHandlerKeepsParticipant(t *testing.T) {
	handler := NewHandler()
	item := handler.Submit("N-1", "participant", nil)
	if item.Participant != "participant" {
		t.Fatal("participant was not retained")
	}
}
