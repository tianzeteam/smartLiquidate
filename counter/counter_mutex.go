package counter

import "sync"

type MutexCounter struct {
	mu     sync.RWMutex
	number uint64
}

func NewMutexCounter(_initNumber uint64) Counter {
	return &MutexCounter{number: _initNumber}
}

func (c *MutexCounter) Add(num uint64) uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.number = c.number + num
	return c.number - 1
}

func (c *MutexCounter) Read() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.number
}
