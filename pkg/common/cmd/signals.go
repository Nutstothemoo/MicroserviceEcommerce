package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Context returns a context that is cancelled when the application receives	

func Context() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigs := make(chan os.Signal, 1)
		defer cancel()
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select{
			case <-sigs:
				cancel()
			case <-ctx.Done():
				return
		}
	}()
	return ctx
}

