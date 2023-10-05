package syncronize

import "sync"

type Counter struct {
	mtx   sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
