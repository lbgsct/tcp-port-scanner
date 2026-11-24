package portscan

import (
	"context"
	"errors"
	"net"
	"sync"
)

type job struct {
	host string
	ip   net.IP
	port uint16
}

func (s *Scanner) run(ctx context.Context, targets []target, ports []uint16, results chan<- Result) {
	defer close(results)
	jobs := make(chan job)
	var workers sync.WaitGroup
	for range s.config.concurrency {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for work := range jobs {
				send(ctx, results, probe(ctx, s.config.timeout, work))
			}
		}()
	}
	for _, target := range targets {
		if ctx.Err() != nil {
			break
		}
		ips, err := resolve(ctx, s.resolver, target)
		if err != nil {
			state := Error
			if errors.Is(ctx.Err(), context.Canceled) {
				state = Canceled
			}
			for _, port := range ports {
				if !send(ctx, results, Result{Host: target.host, Port: port, State: state, Err: err}) {
					break
				}
			}
			continue
		}
		for _, ip := range ips {
			for _, port := range ports {
				select {
				case jobs <- job{target.host, ip, port}:
				case <-ctx.Done():
					close(jobs)
					workers.Wait()
					return
				}
			}
		}
	}
	close(jobs)
	workers.Wait()
}

func send(ctx context.Context, results chan<- Result, result Result) bool {
	select {
	case results <- result:
		return true
	case <-ctx.Done():
		return false
	}
}
