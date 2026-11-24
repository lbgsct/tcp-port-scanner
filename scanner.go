package portscan

import (
	"context"
	"errors"
	"time"
)

type Scanner struct {
	config   config
	resolver resolver
}

func New(options ...Option) (*Scanner, error) {
	c := config{concurrency: 64, timeout: time.Second}
	for _, option := range options {
		if option != nil {
			if err := option(&c); err != nil {
				return nil, err
			}
		}
	}
	return &Scanner{config: c, resolver: systemResolver{}}, nil
}

func (s *Scanner) Scan(ctx context.Context, hosts []string, ports PortSet) (<-chan Result, error) {
	if ctx == nil {
		return nil, errors.New("portscan: nil context")
	}
	targets, err := normalizeHosts(hosts)
	if err != nil {
		return nil, err
	}
	portList, err := ports.expand()
	if err != nil {
		return nil, err
	}
	results := make(chan Result, s.config.concurrency)
	go s.run(ctx, targets, portList, results)
	return results, nil
}
