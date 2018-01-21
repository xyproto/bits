package bits

import (
	"testing"
)

func TestBitsToString(t *testing.T) {
	a := Bits{1, 0, 0, 1, 0, 1, 0}.String()
	b := "1 0 0 1 0 1 0"
	if a != b {
		t.Error("a != b", a, b)
	}
}

func TestChoices(t *testing.T) {
	c := StringToChoices("0 1 | 1 0")
	if c[0][0] != B0 {
		t.Error("first bit of first element should be 0")
	}
	if c[0][1] != B1 {
		t.Error("seond bit of first element should be 1")
	}
	if c[1][0] != B1 {
		t.Error("first bit of second element should be 1")
	}
	if c[1][1] != B0 {
		t.Error("seond bit of second element should be 0")
	}
}
