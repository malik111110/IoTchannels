
package simulator

import (
	"context"
	"sync"
	"chanIoT/pkg/collector"
	"chanIoT/pkg/event"
	"chanIoT/pkg/iot"
)

type Simulator struct {
	Events  chan event.Event
	done    chan bool
	cancel  context.CancelFunc
	wg      *sync.WaitGroup
	Running bool
	mutex   sync.Mutex
}

func New() *Simulator {
	return &Simulator{
		Events:  make(chan event.Event, 100),
		done:    make(chan bool),
		wg:      &sync.WaitGroup{},
		Running: false,
	}
}

func (s *Simulator) Start() {
	s.mutex.Lock()
	if s.Running {
		s.mutex.Unlock()
		return
	}
	s.Running = true
	s.mutex.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	go collector.Collector(s.Events, s.done)

	s.wg.Add(3)
	go iot.SimulateDevice(ctx, "temp-1", event.Temp, s.Events, s.wg)
	go iot.SimulateDevice(ctx, "motion-1", event.Motion, s.Events, s.wg)
	go iot.SimulateDevice(ctx, "door-1", event.Door, s.Events, s.wg)
}

func (s *Simulator) Stop() {
	s.mutex.Lock()
	if !s.Running {
		s.mutex.Unlock()
		return
	}
	s.Running = false
	s.mutex.Unlock()

	if s.cancel != nil {
		s.cancel()
	}
	s.wg.Wait()
	close(s.Events)
	<-s.done
}
