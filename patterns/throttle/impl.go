package throttle

import (
	"fmt"
	"context"
	"time"
)
func ThrottleImpl(ctx context.Context, d time.Duration){
	circuit := func(ctx context.Context) (string, error) {
		// Simulate a circuit operation
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Circuit operation executed for")
		return "Circuit Result", nil
	}

	// Debounce the circuit
	throttle := Throttle(circuit, 1, 1, d)

	throttle(ctx)
	throttle(ctx)
	throttle(ctx)
	throttle(ctx)
	throttle(ctx)
	throttle(ctx)
	throttle(ctx)
}