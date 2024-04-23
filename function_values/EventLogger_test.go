package main

import (
	"testing"
	"time"
)

func TestEventLogger(t *testing.T) {
	// Initialize a slice to store events.
	events := make([]Event, 0)
	// Initialize a map to categorize events by type.
	categories := make(map[string][]Event)
	logEvent := createEventLogger(&events, categories)

	// Log some events
	logEvent("info", "Application started")
	logEvent("error", "Failed to load configuration")
	logEvent("info", "User logged in")
	logEvent("warning", "Disk space low")

	// Test totalEventsByType function
	infoCount := totalEventsByType("info", categories)
	if infoCount != 2 {
		t.Errorf("Expected 2 'info' events, got %d", infoCount)
	}

	errorCount := totalEventsByType("error", categories)
	if errorCount != 1 {
		t.Errorf("Expected 1 'error' event, got %d", errorCount)
	}

	// Test numberOfEventsAfter function
	currentTime := time.Now().UnixMilli()
	eventsAfterCount := numberOfEventsAfter(currentTime-2000, events)
	if eventsAfterCount != 4 {
		t.Errorf("Expected 4 events after a certain timestamp, got %d", eventsAfterCount)
	}

	// Test listEventsByType function
	errorEvents := listEventsByType("error", categories)
	if len(errorEvents) != 1 {
		t.Errorf("Expected 1 'error' event, got %d", len(errorEvents))
	}
	if errorEvents[0].Message != "Failed to load configuration" {
		t.Errorf("Expected 'Failed to load configuration' error message, got %s", errorEvents[0].Message)
	}
}
