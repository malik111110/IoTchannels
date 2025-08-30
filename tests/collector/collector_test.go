package collector_test

import (
	"chanIoT/pkg/collector"
	"chanIoT/pkg/event"
	"testing"
	"time"
)

func TestCollector(t *testing.T) {
	events := make(chan event.Event, 10)
	done := make(chan bool)

	go collector.Collector(events, done)

	// Test case: Door opens, no motion, wait 11 seconds, expect alert
	events <- event.Event{DeviceID: "door-1", Type: event.Door, Value: true, Time: time.Now()}
	time.Sleep(11 * time.Second)

	// TODO: capture stdout and check for alert message

	// Test case: Door opens, motion detected, wait 11 seconds, no alert
	events <- event.Event{DeviceID: "door-1", Type: event.Door, Value: true, Time: time.Now()}
	events <- event.Event{DeviceID: "motion-1", Type: event.Motion, Value: 1, Time: time.Now()}
	time.Sleep(11 * time.Second)

	// TODO: capture stdout and check for no alert message

	close(events)
	<-done
}
