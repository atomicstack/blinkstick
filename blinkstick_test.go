package blinkstick

import "testing"

func TestList_NoPanic(t *testing.T) {
	got := List()
	if got == nil {
		t.Fatal("List() returned nil; expected empty slice when no device present")
	}
}
