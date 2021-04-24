package util

import (
	"context"
	"time"
)

// Sleep is similar to time.Sleep but it is context aware.
func Sleep(ctx context.Context, d time.Duration) bool {
	t := time.NewTimer(d)

	select {
	case <-t.C:
		return true
	case <-ctx.Done():
		t.Stop()
		return false
	}
}
