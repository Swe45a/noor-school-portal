package web

import (
	"net/http"
	"sync"
	"time"

	apperrors "school-portal/internal/errors"
	"school-portal/internal/utils"
)

type visitor struct {
	count       int
	windowStart time.Time
}

// IPRateLimiter is a simple fixed-window, per-IP request limiter. It is in-memory and
// per-instance, which is sufficient for the school's single-instance deployment; it is not
// meant to survive a multi-node rollout without moving the counters to a shared store.
type IPRateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	limit    int
	window   time.Duration
}

func NewIPRateLimiter(limit int, window time.Duration) *IPRateLimiter {
	rl := &IPRateLimiter{
		visitors: make(map[string]*visitor),
		limit:    limit,
		window:   window,
	}
	go rl.cleanupLoop()
	return rl
}

func (rl *IPRateLimiter) cleanupLoop() {
	ticker := time.NewTicker(rl.window * 2)
	defer ticker.Stop()
	for range ticker.C {
		cutoff := time.Now().Add(-rl.window * 2)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if v.windowStart.Before(cutoff) {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *IPRateLimiter) allow(ip string) bool {
	now := time.Now()

	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, ok := rl.visitors[ip]
	if !ok || now.Sub(v.windowStart) > rl.window {
		rl.visitors[ip] = &visitor{count: 1, windowStart: now}
		return true
	}

	if v.count >= rl.limit {
		return false
	}

	v.count++
	return true
}

// Middleware rejects requests over the configured limit with 429 rate_limited, keyed by client IP.
func (rl *IPRateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := utils.ClientIP(r)
		if !rl.allow(ip) {
			utils.WriteError(w, apperrors.ErrRateLimited)
			return
		}
		next.ServeHTTP(w, r)
	})
}
