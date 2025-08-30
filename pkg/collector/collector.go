package collector

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"chanIoT/pkg/event"
)

var lastMotionTime time.Time
var doorOpenTime time.Time
var isDoorOpen bool

func Collector(in <-chan event.Event, done chan<- bool) {
	file, err := os.OpenFile("events.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	for evt := range in {
		fmt.Printf("[%s] %s -> %v at %s\n",
			evt.DeviceID, evt.Type.String(), evt.Value, evt.Time.Format(time.RFC3339))

		if err := encoder.Encode(evt); err != nil {
			fmt.Println("Error encoding JSON:", err)
		}

		switch evt.Type {
		case event.Door:
			isDoorOpen = evt.Value.(bool)
			if isDoorOpen {
				doorOpenTime = evt.Time
			} else {
				doorOpenTime = time.Time{}
			}
		case event.Motion:
			if evt.Value.(int) == 1 {
				lastMotionTime = evt.Time
			}
		}

		if isDoorOpen && time.Since(doorOpenTime) > 10*time.Second && time.Since(lastMotionTime) > 10*time.Second {
			fmt.Println("ALERT: Door open for more than 10 seconds with no motion!")
		}
	}
	done <- true
}
