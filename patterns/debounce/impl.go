package debounce

import (
	"context"
	"fmt"
	"time"
)

func DebounceFirstImpl(ctx context.Context, d time.Duration) {
	circuit := func(ctx context.Context, idx int) (string, error) {
		// Simulate a circuit operation
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Circuit operation executed for", idx)
		return "Circuit Result", nil
	}

	// Debounce the circuit
	debouncedCircuit := DebounceFirst(circuit, d)

	for i := 1; i <= 5; i++ {
		debouncedCircuit(ctx, i)
	}
}

func DebounceLastImpl(ctx context.Context, d time.Duration) {
	circuit := func(ctx context.Context, idx int) (string, error) {
		// Simulate a circuit operation
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Circuit operation executed for", idx)
		return "Circuit Result", nil
	}

	// Debounce the circuit
	debouncedCircuit := DebounceFirst(circuit, d)

	debouncedCircuit(ctx, 1)
	debouncedCircuit(ctx, 2)
	debouncedCircuit(ctx, 3)
	debouncedCircuit(ctx, 4)
	debouncedCircuit(ctx, 5)
	debouncedCircuit(ctx, 6)
	debouncedCircuit(ctx, 7) // Last call will execute
}