package rateLimiter

import (
	"context"
	"log"
	"os"
	"sync"
	"time"
)

type IAPIConnection interface {
	ReadFile(context.Context) error
	ResolveAddress(context.Context) error
	DemoFunc()
}
type APIConnection struct {
	conn IAPIConnection
}

func (a *APIConnection) ReadFile(context.Context) error {
	return nil
}
func (a *APIConnection) ResolveAddress(context.Context) error {
	return nil
}

func (a *APIConnection) DemoFunc() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	// log.SetOutput(ioutil.Discard)
	log.SetFlags(log.Ltime | log.LUTC)

	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := a.conn.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := a.conn.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}

func OpenNoRateLimit() *NoRateLimitAPIConnection {
	return &NoRateLimitAPIConnection{}
}

type NoRateLimitAPIConnection struct {
	APIConnection
}

func (a *NoRateLimitAPIConnection) ReadFile(ctx context.Context) error {
	// Pretend we do work here
	time.Sleep(300 * time.Microsecond)
	return nil
}

func (a *NoRateLimitAPIConnection) ResolveAddress(ctx context.Context) error {
	// Pretend we do work here
	time.Sleep(300 * time.Microsecond)
	return nil
}
