package rateLimiter

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func OpenRateLimit() *RateLimitAPIConnection {
	return &RateLimitAPIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1000000), 50000),
	}
}

type RateLimitAPIConnection struct {
	APIConnection
	rateLimiter *rate.Limiter
}

func (a *RateLimitAPIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil { // <2>
		fmt.Println(err)
		return err
	}
	// Pretend we do work here
	time.Sleep(300 * time.Microsecond)
	return nil
}

func (a *RateLimitAPIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil { // <2>
		fmt.Println(err)
		return err
	}
	// Pretend we do work here
	time.Sleep(300 * time.Microsecond)
	return nil
}
