package portscan

import (
	"context"
	"syscall"
	"testing"
)

func TestDialState(t *testing.T) {
	if got := dialState(context.Background(), syscall.ECONNREFUSED); got != Closed {
		t.Fatalf("got %s", got)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := dialState(ctx, context.Canceled); got != Canceled {
		t.Fatalf("got %s", got)
	}
}
