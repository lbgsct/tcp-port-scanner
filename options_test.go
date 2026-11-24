package portscan

import (
	"errors"
	"testing"
	"time"
)

func TestOptions(t *testing.T) {
	scanner, err := New(WithConcurrency(3), WithConnectTimeout(20*time.Millisecond))
	if err != nil || scanner.config.concurrency != 3 || scanner.config.timeout != 20*time.Millisecond {
		t.Fatalf("unexpected scanner: %#v, %v", scanner, err)
	}
	if _, err := New(WithConcurrency(0)); !errors.Is(err, ErrInvalidConcurrency) {
		t.Fatalf("got %v", err)
	}
	if _, err := New(WithConnectTimeout(0)); !errors.Is(err, ErrInvalidTimeout) {
		t.Fatalf("got %v", err)
	}
}
