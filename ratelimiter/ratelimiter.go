package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mut             sync.Mutex
	rateLimits      map[time.Duration]int
	rateLimitStarts map[time.Duration]time.Time
	rateLimitCounts map[time.Duration]int
}

func NewRateLimiter(rateLimits map[time.Duration]int) *RateLimiter {
	return &RateLimiter{
		rateLimits:      rateLimits,
		rateLimitStarts: make(map[time.Duration]time.Time),
		rateLimitCounts: make(map[time.Duration]int),
	}
}

func (rl *RateLimiter) Wait() bool {
	rl.mut.Lock()
	defer rl.mut.Unlock()

	delay := rl.getDelay()
	if delay > 0 {
		return false
	}
	rl.updateDelay()
	return true
}

func (rl *RateLimiter) WaitRetries(retries int) bool {
	for i := 0; i < retries; i++ {
		if rl.Wait() {
			return true
		}
		time.Sleep(time.Millisecond * 300)
	}
	return false
}

func (rl *RateLimiter) getDelay() time.Duration {
	now := time.Now()

	var delay time.Duration

	for timeSpan, count := range rl.rateLimitCounts {
		limit := rl.rateLimits[timeSpan]
		if count >= limit {
			start := rl.rateLimitStarts[timeSpan]
			newDelay := start.Add(timeSpan).Sub(now)
			if newDelay > delay {
				delay = newDelay
			}
		}
	}
	return delay
}

func (rl *RateLimiter) updateDelay() {
	now := time.Now()

	for timeSpan := range rl.rateLimits {
		var count int
		start, exists := rl.rateLimitStarts[timeSpan]

		if !exists || start.Before(now.Add(-timeSpan)) {
			rl.rateLimitStarts[timeSpan] = now
		} else {
			count = rl.rateLimitCounts[timeSpan]
		}

		rl.rateLimitCounts[timeSpan] = count + 1
	}
}
