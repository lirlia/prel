package testutil

import (
	"sync"
	"time"
)

type Clock struct {
	mu  *sync.RWMutex
	now time.Time
}

func NewClock() *Clock {
	return &Clock{
		mu:  &sync.RWMutex{},
		now: time.Now(),
	}
}

func (c *Clock) Now() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.now
}

func (c *Clock) Set(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = t
}

func (h *TestHelper) SetClock(t time.Time) {
	h.Clock.Set(t)
}

// Travel は、 Clock の時刻を t に移動して関数 fn を呼んだあと、時刻をもとに戻します。
func (h *TestHelper) Travel(t time.Time, fn func()) {
	now := h.Clock.Now()
	h.Clock.Set(t)
	defer h.Clock.Set(now)
	fn()
}
