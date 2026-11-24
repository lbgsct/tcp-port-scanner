package portscan

import (
	"errors"
	"testing"
)

func TestErrorsAreSentinels(t *testing.T) {
	if !errors.Is(ErrInvalidHost, ErrInvalidHost) || ErrInvalidHost.Error() == "" {
		t.Fatal("invalid sentinel error")
	}
}
