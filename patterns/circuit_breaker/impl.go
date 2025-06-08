package ctbreaker

import (
	"context"
)

func CtBreaker(ctx context.Context) {
	circuit := getRandomCircuit

	circuitBreaker := Breaker(circuit, 2)

	for i := 0; i < 5; i++ {
		_, err := circuitBreaker(ctx)
		if err != nil {
			println(err.Error())
		}
	}
}