package counter

import (
	"fmt"
	"time"
)

// Counter prints itself value to stdout every second
type Counter struct {
	value int

	cancelCh chan struct{}
	doneCh   chan struct{}
}

// Start the counter
func (c *Counter) Start() {
	go func() {
		timer := time.NewTicker(time.Second)
		for {
			select {
			case <-timer.C:
				c.value++
				fmt.Printf("Application has been working for %d seconds (press Ctrl+C to exit)\n", c.value)
			case <-c.cancelCh:
				fmt.Println("Counter stopped...")
				c.doneCh <- struct{}{}
				return
			}
		}
	}()

	fmt.Println("Counter started...")
}

// Stop the counter
func (c *Counter) Stop() chan struct{} {
	c.cancelCh <- struct{}{}
	return c.doneCh
}

// New creates a Counter
func New() Counter {
	return Counter{
		cancelCh: make(chan struct{}, 1),
		doneCh:   make(chan struct{}, 1),
	}
}
