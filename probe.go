package portscan

import (
	"context"
	"errors"
	"fmt"
	"net"
	"syscall"
	"time"
)

func probe(ctx context.Context, timeout time.Duration, work job) Result {
	started := time.Now()
	dialCtx, cancel := context.WithTimeout(ctx, timeout)
	connection, err := (&net.Dialer{}).DialContext(dialCtx, "tcp", net.JoinHostPort(work.ip.String(), fmt.Sprintf("%d", work.port)))
	cancel()
	if connection != nil {
		_ = connection.Close()
	}
	return Result{Host: work.host, IP: append(net.IP(nil), work.ip...), Port: work.port, State: dialState(ctx, err), Duration: time.Since(started), Err: err}
}

func dialState(ctx context.Context, err error) State {
	if err == nil {
		return Open
	}
	if errors.Is(ctx.Err(), context.Canceled) {
		return Canceled
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return Timeout
	}
	var networkErr net.Error
	if errors.As(err, &networkErr) && networkErr.Timeout() {
		return Timeout
	}
	if errors.Is(err, syscall.ECONNREFUSED) {
		return Closed
	}
	if errors.Is(err, syscall.EHOSTUNREACH) || errors.Is(err, syscall.ENETUNREACH) {
		return Unreachable
	}
	return Error
}
