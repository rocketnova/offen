package persistence

import (
	"strings"
	"testing"
	"time"
)

func TestNewEventID(t *testing.T) {
	first, firstErr := NewEventID()
	if firstErr != nil {
		t.Fatalf("Unexpected error %v", firstErr)
	}

	time.Sleep(time.Millisecond)

	second, secondErr := NewEventID()
	if secondErr != nil {
		t.Fatalf("Unexpected error %v", secondErr)
	}

	if strings.Compare(first, second) != -1 {
		t.Errorf("Expected first event id to sort lower, got %s and %s", first, second)
	}
}
