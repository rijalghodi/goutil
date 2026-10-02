package validate

import "testing"

func TestPositive(t *testing.T) {
	if err := Positive(1, "size"); err != nil {
		t.Errorf("Positive(1) unexpected error: %v", err)
	}
	if err := Positive(0, "size"); err == nil {
		t.Error("Positive(0) expected error")
	}
	if err := Positive(-3, "size"); err == nil {
		t.Error("Positive(-3) expected error")
	}
}
