package sample

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

func OpenRateLimit() *RateLimitAPIConnection {
	return &RateLimitAPIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

type RateLimitAPIConnection struct {
	APIConnection
	rateLimiter *rate.Limiter
}

func (a *RateLimitAPIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil { // <2>
		return err
	}
	// Pretend we do work here
	time.Sleep(100 * time.Microsecond)
	return nil
}

func (a *RateLimitAPIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil { // <2>
		return err
	}
	// Pretend we do work here
	time.Sleep(100 * time.Microsecond)
	return nil
}
