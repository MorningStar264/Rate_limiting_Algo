package tokenbucket

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity     float64
	tokens       float64
	refillRate   float64
	lastRefilled time.Time
	mu           sync.Mutex
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		lastRefilled: time.Now(),
	}
}

func (tb *TokenBucket) AllowRequest(tokensNeeded float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	now := time.Now()
	elpased := now.Sub(tb.lastRefilled).Seconds()
	tb.tokens += elpased * tb.refillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	if tb.tokens >= tokensNeeded {
		tb.tokens -= tokensNeeded
		tb.lastRefilled = now
		return true
	}
	tb.lastRefilled = now
	return false
}
