package portscan

import "time"

type Option func(*config) error

type config struct {
	concurrency int
	timeout     time.Duration
}

func WithConcurrency(n int) Option {
	return func(c *config) error {
		if n <= 0 {
			return ErrInvalidConcurrency
		}
		c.concurrency = n
		return nil
	}
}

func WithConnectTimeout(timeout time.Duration) Option {
	return func(c *config) error {
		if timeout <= 0 {
			return ErrInvalidTimeout
		}
		c.timeout = timeout
		return nil
	}
}
