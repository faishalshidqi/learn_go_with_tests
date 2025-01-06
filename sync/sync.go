package sync

import "sync"

type CounterMethods interface {
	Inc()
	Value() int
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
func (c *Counter) Value() int {
	return c.value
}
