package util

import "context"

// Block blocks the calling goroutine forever.
func Block() {
	select {} // (or) <-make(<-chan struct{})
}

// BlockWithContext is context aware Block.
func BlockWithContext(ctx context.Context) {
	<-ctx.Done()
}
