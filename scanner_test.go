package portscan

import (
	"context"
	"errors"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestScanOpenAndClosed(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if errors.Is(err, syscall.EPERM) {
		t.Skip("sandbox does not permit listening sockets")
	}
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	open := listener.Addr().(*net.TCPAddr).Port
	other, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	closed := other.Addr().(*net.TCPAddr).Port
	other.Close()
	scanner, _ := New(WithConcurrency(2), WithConnectTimeout(time.Second))
	results, err := scanner.Scan(context.Background(), []string{"127.0.0.1"}, Ports(open, closed))
	if err != nil {
		t.Fatal(err)
	}
	got := map[int]State{}
	for result := range results {
		got[int(result.Port)] = result.State
	}
	if got[open] != Open || got[closed] != Closed {
		t.Fatalf("got %v", got)
	}
}

func TestScanCanceledBeforeStart(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	scanner, _ := New()
	results, err := scanner.Scan(ctx, []string{"127.0.0.1"}, Ports(80))
	if err != nil {
		t.Fatal(err)
	}
	for range results {
		t.Fatal("unexpected result")
	}
}
