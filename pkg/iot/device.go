package iot

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"chanIoT/pkg/event"
)

func SimulateDevice(ctx context.Context, id string, typ event.DeviceType, out chan<- event.Event, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			var value interface{}

			switch typ {
			case event.Temp:
				value = 18 + rand.Intn(13) // 18–30°C
				time.Sleep(2 * time.Second)
			case event.Motion:
				value = rand.Intn(2) // 0 ou 1
				time.Sleep(3 * time.Second)
			case event.Door:
				value = rand.Intn(2) == 1 // true/false
				time.Sleep(3 * time.Second)
			}

			evt := event.Event{
				DeviceID: id,
				Type:     typ,
				Value:    value,
				Time:     time.Now(),
			}

			out <- evt
		}
	}
}
