package rate_limiter

import "sync"

type limiter struct {
	countReqs int
	tokens    int
	mu        sync.Mutex
}

func NewLimiter(countReqs int) *limiter {
	return &limiter{
		countReqs: countReqs,
		tokens:    countReqs,
	}
}

func (l *limiter) CanDoWork() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.tokens > 0 {
		l.tokens = l.tokens - 1
		return true
	}

	return false
}

func (l *limiter) Done() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.tokens += 1
	if l.tokens > l.countReqs {
		l.tokens = l.countReqs
	}
}
