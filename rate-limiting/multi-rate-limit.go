package rateLimiter

import (
	"context"
	"sort"
	"time"

	"golang.org/x/time/rate"
)

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

func OpenMultiRateLimit() *MultiRateLimitAPIConnection {
	secondLimit := rate.NewLimiter(Per(2, time.Second), 1)   // <1>
	minuteLimit := rate.NewLimiter(Per(10, time.Minute), 10) // <2>
	return &MultiRateLimitAPIConnection{
		rateLimiter: MultiLimiter(secondLimit, minuteLimit), // <3>
	}
}

type MultiRateLimitAPIConnection struct {
	APIConnection
	rateLimiter RateLimiter
}

func (a *MultiRateLimitAPIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// Pretend we do work here
	// time.Sleep(100 * time.Microsecond)
	return nil
}

func (a *MultiRateLimitAPIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// Pretend we do work here
	// time.Sleep(100 * time.Microsecond)
	return nil
}

type RateLimiter interface { // <1>
	Wait(context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit) // <2>
	return &multiLimiter{limiters: limiters}
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit() // <3>
}
