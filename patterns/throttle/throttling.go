package throttle

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Effector func(context.Context) (string, error)

func Throttle(e Effector, max uint, refill uint, d time.Duration) Effector {
	var (
		tokens uint = max
		once   sync.Once
	)

	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}
		once.Do(func() {
			ticker := time.NewTicker(d)

			go func() {
				defer ticker.Stop()

				for {
					select {
					case <- ctx.Done():
						return
					
					case <- ticker.C:
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})
		if tokens <= 0 {
			return "", fmt.Errorf("Too many calls")
		}
		tokens--
		return e(ctx)
	}
}
