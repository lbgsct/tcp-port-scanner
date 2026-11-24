package portscan

import (
	"errors"
	"reflect"
	"testing"
)

func TestPortSetExpansion(t *testing.T) {
	ports, err := Combine(Ports(443, 80, 80), Range(81, 82)).expand()
	if err != nil || !reflect.DeepEqual(ports, []uint16{80, 81, 82, 443}) {
		t.Fatalf("got %v, %v", ports, err)
	}
	if _, err := Ports(0).expand(); !errors.Is(err, ErrInvalidPort) {
		t.Fatalf("got %v", err)
	}
	if _, err := Range(2, 1).expand(); !errors.Is(err, ErrInvalidRange) {
		t.Fatalf("got %v", err)
	}
}
