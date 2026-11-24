package portscan_test

import (
	"context"
	"log"
	"time"

	"portscan"
)

func ExampleScanner_Scan() {
	scanner, err := portscan.New(portscan.WithConcurrency(20), portscan.WithConnectTimeout(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	results, err := scanner.Scan(context.Background(), []string{"127.0.0.1"}, portscan.Combine(portscan.Ports(22), portscan.Range(80, 81)))
	if err != nil {
		log.Fatal(err)
	}
	for result := range results {
		_ = result
	}
}
