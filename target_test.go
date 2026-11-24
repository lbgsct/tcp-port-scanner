package portscan

import (
	"errors"
	"testing"
)

func TestNormalizeHosts(t *testing.T) {
	targets, err := normalizeHosts([]string{"127.0.0.1", "127.0.0.1", "EXAMPLE.org", "example.org."})
	if err != nil || len(targets) != 2 || !targets[0].isLiteral {
		t.Fatalf("got %#v, %v", targets, err)
	}
	for _, host := range []string{"", "300.1.1.1", "bad host", "-bad.example"} {
		if _, err := normalizeHosts([]string{host}); !errors.Is(err, ErrInvalidHost) {
			t.Fatalf("%q: %v", host, err)
		}
	}
}
