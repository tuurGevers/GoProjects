package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/elliotchance/pie/v2"
)

// Event represents an event in the log.
type Event struct {
	Type      string
	Message   string
	Timestamp int64
}

// createEventLogger returns a closure that logs events and provides statistics.
func createEventLogger(events *[]Event, categories map[string][]Event) func(eventType, message string) int {

	// The returned closure that handles logging.
	return func(eventType, message string) int {
		// TODO: Create a new event with the current timestamp.
		NewEvent := Event{
			eventType,
			message,
			time.Now().UnixMilli(),
		}
		// TODO: Append the event to the slice of all events.
		*events = append(*events, NewEvent)
		// TODO: Add the event to the correct category in the map.
		categories[eventType] = append(categories[eventType], NewEvent)
		// TODO: Return the total number of events logged so far.
		return len(*events)
	}
}

// totalEventsByType returns the total number of events of a given type.
func totalEventsByType(eventType string, categories map[string][]Event) int {
	// TODO: Return the count of events in the category specified by eventType.
	return len(categories[eventType])
}

// listEventsByType returns a slice of all events of a given type.
func listEventsByType(eventType string, categories map[string][]Event) (sortedEvents []Event) {
	// TODO: Return all events categorized under eventType.
	sortedEvents = categories[eventType]
	sort.Slice(sortedEvents, func(i, j int) bool {
		return sortedEvents[i].Timestamp < sortedEvents[j].Timestamp
	})

	return
}

// numberOfEventsAfter returns the count of all events that occurred after the given timestamp.
func numberOfEventsAfter(timestamp int64, events []Event) int {
	// TODO: Iterate through all events and count those with a timestamp greater than the provided value.
	fmt.Println(len(events))
	return len(pie.Filter(events, func(e Event) bool {
		fmt.Print("event ts")
		printReadableTime(e.Timestamp)

		fmt.Print("ts")
		printReadableTime(timestamp)

		return e.Timestamp > timestamp
	}))

}

func printReadableTime(ts int64) {
	// Convert Unix millisecond timestamp to seconds and nanoseconds for Go's time package
	t := time.Unix(ts/1000, (ts%1000)*1000000)
	// Format time in a readable format
	fmt.Println(t.UTC().Format("2006-01-02 15:04:05"))
}

func mainLogger() {
	// TODO: Initialize a slice to store events.
	events := make([]Event, 0)
	// TODO: Initialize a map to categorize events by type.
	categories := make(map[string][]Event)
	// Example usage of the logger and statistics functions.
	logEvent := createEventLogger(&events, categories)

	printReadableTime(time.Now().UnixMilli())

	logEvent("info", "Application started")
	time.Sleep(4 * time.Second) // Sleep for 4 seconds

	printReadableTime(time.Now().UnixMilli())

	logEvent("error", "Failed to load configuration")

	// Add more logEvent calls to simulate activity.

	fmt.Println("Total 'info' events:", totalEventsByType("info", categories))
	fmt.Println("Total 'error' events:", totalEventsByType("error", categories))

	fmt.Println("Events after a certain timestamp:", numberOfEventsAfter(time.Now().UnixMilli()-2000, events))
}
