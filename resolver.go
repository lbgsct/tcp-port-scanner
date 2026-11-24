package portscan

import (
	"context"
	"errors"
	"net"
	"net/netip"
)

type resolver interface {
	LookupNetIP(context.Context, string, string) ([]netip.Addr, error)
}
type systemResolver struct{}

func (systemResolver) LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error) {
	return net.DefaultResolver.LookupNetIP(ctx, network, host)
}

func resolve(ctx context.Context, resolver resolver, target target) ([]net.IP, error) {
	if target.isLiteral {
		return []net.IP{net.IP(target.literal.AsSlice())}, nil
	}
	addresses, err := resolver.LookupNetIP(ctx, "ip", target.host)
	if err != nil {
		return nil, err
	}
	seen := make(map[netip.Addr]struct{}, len(addresses))
	ips := make([]net.IP, 0, len(addresses))
	for _, address := range addresses {
		address = address.Unmap()
		if _, ok := seen[address]; !ok {
			seen[address] = struct{}{}
			ips = append(ips, net.IP(address.AsSlice()))
		}
	}
	if len(ips) == 0 {
		return nil, errors.New("portscan: DNS name returned no IP addresses")
	}
	return ips, nil
}
