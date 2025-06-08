package retry

import (
	"context"
	"fmt"
)

var count int

func EmulateTransientError(ctx context.Context) (string, error) {
	// Simulate a transient error
	if count <= 3 {
		return "intentional fail", fmt.Errorf("transient error")
	} else {
		return "success", nil
	}
}