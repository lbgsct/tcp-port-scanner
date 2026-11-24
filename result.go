package portscan

import (
	"net"
	"time"
)

type State uint8

const (
	Open State = iota
	Closed
	Timeout
	Unreachable
	Canceled
	Error
)

func (s State) String() string {
	switch s {
	case Open:
		return "open"
	case Closed:
		return "closed"
	case Timeout:
		return "timeout"
	case Unreachable:
		return "unreachable"
	case Canceled:
		return "canceled"
	default:
		return "error"
	}
}

type Result struct {
	Host     string
	IP       net.IP
	Port     uint16
	State    State
	Duration time.Duration
	Err      error
}
