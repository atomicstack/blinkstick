package blinkstick

import (
	"strings"
	"testing"
)

func TestGetColor_LowercaseError(t *testing.T) {
	_, err := GetColor("not-a-real-color", 50)
	if err == nil {
		t.Fatal("expected error for invalid colour, got nil")
	}
	if msg := err.Error(); strings.ToLower(msg[:1]) != msg[:1] {
		t.Errorf("error message must start lowercase: %q", msg)
	}
}
