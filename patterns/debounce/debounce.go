package debounce

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Circuit func(ctx context.Context, idx int) (string, error)

func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var (
		threshold time.Time
		result    string
		err       error
		m         sync.Mutex
	)

	return func(ctx context.Context, idx int) (string, error) {
		m.Lock()

		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()
		if time.Now().Before(threshold) {
			return result, err
		}
		result, err = circuit(ctx, idx)
		return result, err
	}
}

func DebounceLast(circuit Circuit, d time.Duration) Circuit {
	var (
		threshold time.Time
		ticker    *time.Ticker
		result    string
		err       error
		once      sync.Once
		m         sync.Mutex
	)

	return func(ctx context.Context, idx int) (string, error) {
		fmt.Printf("DebounceLast: %d\n", idx)
		m.Lock()
		defer m.Unlock()

		threshold = time.Now().Add(d)

		once.Do(func() {
			ticker = time.NewTicker(time.Millisecond * 100)
			go func() {
				defer func() {
					m.Lock()
					ticker.Stop()
					once = sync.Once{}
					m.Unlock()
				}()

				for {
					select {
					case <-ticker.C:
						m.Lock()
						if time.Now().After(threshold) {
							result, err = circuit(ctx, idx)
							m.Unlock()
							return
						}

						m.Unlock()
					case <-ctx.Done():
						m.Lock()
						result, err = "", ctx.Err()
						m.Unlock()
						return
					}
				}
			}()
		})
		return result, err
	}
}
