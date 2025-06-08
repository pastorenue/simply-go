package ctbreaker

import (
	"fmt"
	"context"
	"errors"
	"math/rand"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var (
		consecutiveFailures int = 0
		lastAttempt             = time.Now()
		m                   sync.RWMutex
	)
	return func(ctx context.Context) (string, error) {
		m.RLock()

		d := consecutiveFailures - int(failureThreshold)
		if d >= 0 {
			shoulldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shoulldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unavailable")

			}
		}

		m.RUnlock()
		response, err := circuit(ctx)
		m.Lock()

		defer m.Unlock()
		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++
			return response, err
		}
		consecutiveFailures = 0
		return response, nil
	}
}

func getRandomCircuit(ctx context.Context) (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, value int) {
		fmt.Fprintf(w, "%s\t%d\n", name, value)
	}
	show("Random Number", r.Intn(10))
	if r.Intn(10) == 5 {
		return "", errors.New("random error")
	}
	return "random success", nil
}