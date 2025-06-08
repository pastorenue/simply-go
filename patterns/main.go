package main

import (
	"context"
	"time"

	"github.com/pastorenue/patterns/throttle"
)

func main() {
	ctx := context.Background()
	// Call the circuit breaker function
	throttle.ThrottleImpl(ctx, time.Millisecond * 1)
}
