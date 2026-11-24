package portscan

import (
	"context"
	"errors"
	"net/netip"
	"testing"
)

type fakeResolver struct {
	addresses []netip.Addr
	err       error
}

func (r fakeResolver) LookupNetIP(context.Context, string, string) ([]netip.Addr, error) {
	return r.addresses, r.err
}

func TestResolveDeduplicatesAddresses(t *testing.T) {
	addresses := []netip.Addr{netip.MustParseAddr("127.0.0.1"), netip.MustParseAddr("127.0.0.1"), netip.MustParseAddr("::1")}
	ips, err := resolve(context.Background(), fakeResolver{addresses: addresses}, target{host: "example.org"})
	if err != nil || len(ips) != 2 {
		t.Fatalf("got %v, %v", ips, err)
	}
	_, err = resolve(context.Background(), fakeResolver{err: errors.New("no DNS")}, target{host: "example.org"})
	if err == nil {
		t.Fatal("expected resolver error")
	}
}
