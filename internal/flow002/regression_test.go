package flow002

import "testing"

func Test670BusinessRegression(t *testing.T) {
	handler := NewHandler()
	first := handler.Submit("A-001", "张一", []string{"first"})
	second := handler.Submit("A-002", "张二", []string{"second"})
	if first.Number != "A-001" {
		t.Fatalf("first number changed to %s", first.Number)
	}
	if second.Number != "A-002" {
		t.Fatalf("second number is %s", second.Number)
	}
	if second.Participant != "张二" {
		t.Fatalf("second participant is %s", second.Participant)
	}
	if len(second.Tags) != 4 || second.Tags[3] != "second" {
		t.Fatalf("unexpected second tags: %v", second.Tags)
	}
}

func TestHandlerDefaults(t *testing.T) {
	handler := NewHandler()
	item := handler.Prepare("", "")
	if item.Number != "unassigned" || item.Participant != "unknown" {
		t.Fatal("defaults were not applied")
	}
}

func TestHandlerSnapshot(t *testing.T) {
	handler := NewHandler()
	item := handler.Submit("A-001", "张一", []string{"x"})
	snapshot := handler.Snapshot(item)
	if len(snapshot.Tags) != len(item.Tags) {
		t.Fatal("snapshot length mismatch")
	}
	snapshot.Tags[0] = "changed"
	if item.Tags[0] == "changed" {
		t.Fatal("snapshot shares tag storage")
	}
}
