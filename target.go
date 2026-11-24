package portscan

import (
	"fmt"
	"net/netip"
	"strings"
)

type target struct {
	host      string
	literal   netip.Addr
	isLiteral bool
}

func normalizeHosts(hosts []string) ([]target, error) {
	if len(hosts) == 0 {
		return nil, ErrNoHosts
	}
	seen := make(map[string]struct{}, len(hosts))
	targets := make([]target, 0, len(hosts))
	for _, raw := range hosts {
		host := strings.TrimSpace(raw)
		if !validHost(host) {
			return nil, fmt.Errorf("%w: %q", ErrInvalidHost, raw)
		}
		if address, err := netip.ParseAddr(host); err == nil {
			address = address.Unmap()
			key := "ip:" + address.String()
			if _, exists := seen[key]; !exists {
				seen[key] = struct{}{}
				targets = append(targets, target{host: host, literal: address, isLiteral: true})
			}
			continue
		}
		key := "dns:" + strings.ToLower(strings.TrimSuffix(host, "."))
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			targets = append(targets, target{host: host})
		}
	}
	return targets, nil
}

func validHost(host string) bool {
	if host == "" || len(host) > 253 || strings.ContainsAny(host, " \t\r\n/:[]") {
		return false
	}
	if _, err := netip.ParseAddr(host); err == nil {
		return true
	}
	if looksLikeIPv4(host) {
		return false
	}
	for _, label := range strings.Split(strings.TrimSuffix(host, "."), ".") {
		if label == "" || len(label) > 63 || label[0] == '-' || label[len(label)-1] == '-' {
			return false
		}
		for _, char := range label {
			if !(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' || char == '-') {
				return false
			}
		}
	}
	return true
}

func looksLikeIPv4(host string) bool {
	if !strings.Contains(host, ".") {
		return false
	}
	for _, char := range host {
		if char != '.' && (char < '0' || char > '9') {
			return false
		}
	}
	return true
}
