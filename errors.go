package portscan

import "errors"

var (
	ErrInvalidConcurrency = errors.New("portscan: concurrency must be greater than zero")
	ErrInvalidTimeout     = errors.New("portscan: connect timeout must be greater than zero")
	ErrNoHosts            = errors.New("portscan: no hosts")
	ErrInvalidHost        = errors.New("portscan: invalid host")
	ErrNoPorts            = errors.New("portscan: no ports")
	ErrInvalidPort        = errors.New("portscan: port must be between 1 and 65535")
	ErrInvalidRange       = errors.New("portscan: range start is greater than range end")
)
