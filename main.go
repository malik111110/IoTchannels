package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DeviceType int

const (
	temp DeviceType = iota
	motion
	door
)

func (d DeviceType) String() string {
	switch d {
	case temp:
		return "temperature sensor"
	case motion:
		return "motion sensor"
	case door:
		return "smart door"
	default:
		return "Unknown"
	}
}

type Event struct {
	DeviceID string
	Type     DeviceType
	Value    interface{}
	Time     time.Time
}

func simulateDevice(id string, typ DeviceType, out chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		var value interface{}

		switch typ {
		case temp:
			value = 18 + rand.Intn(13) // 18–30°C
			time.Sleep(2 * time.Second)
		case motion:
			value = rand.Intn(2) // 0 ou 1
			time.Sleep(3 * time.Second)
		case door:
			value = rand.Intn(2) == 1 // true/false
			time.Sleep(3 * time.Second)
		}

		evt := Event{
			DeviceID: id,
			Type:     typ,
			Value:    value,
			Time:     time.Now(),
		}

		out <- evt
	}
}

func collector(in <-chan Event, done chan<- bool) {
	for evt := range in {
		fmt.Printf("[%s] %s -> %v at %s\n",
			evt.DeviceID, evt.Type.String(), evt.Value, evt.Time.Format(time.RFC3339))
	}
	done <- true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	events := make(chan Event, 10)
	done := make(chan bool)
	var wg sync.WaitGroup
	go collector(events, done)
	wg.Add(3)
	go simulateDevice("temp-1", temp, events, &wg)
	go simulateDevice("motion-1", motion, events, &wg)
	go simulateDevice("door-1", door, events, &wg)
	wg.Wait()
	close(events)
	<-done
}
