package bits

import (
	"strings"
	"testing"
)

var and = &TruthTable{
	"0 0 -> 0",
	"0 1 -> 0",
	"1 0 -> 0",
	"1 1 -> 1",
}

var andInverted = &ProbTable{
	"0 -> 0 0 | 0 1 | 1 0",
	"1 -> 1 1",
}

var or = &TruthTable{
	"0 0 -> 0",
	"0 1 -> 1",
	"1 0 -> 1",
	"1 1 -> 1",
}

var orInverted = &ProbTable{
	"0 -> 0 0",
	"1 -> 0 1 | 1 0 | 1 1",
}

var xor = &TruthTable{
	"0 0 -> 0",
	"0 1 -> 1",
	"1 0 -> 1",
	"1 1 -> 0",
}

var xorInverted = &ProbTable{
	"0 -> 0 0 | 1 1",
	"1 -> 0 1 | 1 0",
}

func TestTruthTableToGate(t *testing.T) {
	gate := and.Gate()
	if gate(Bits{1, 1}) != 1 {
		t.Error("1 and 1 != 1")
	}
	if gate(Bits{0, 0}) != 0 {
		t.Error("0 and 0 != 0")
	}
	if gate(Bits{0, 1}) != 0 {
		t.Error("0 and 1 != 0")
	}
	if gate(Bits{1, 0}) != 0 {
		t.Error("1 and 0 != 0")
	}
}

func TestComplete(t *testing.T) {
	if !and.Complete() {
		t.Error("and is supposed to be complete")
	}
	if !or.Complete() {
		t.Error("or is supposed to be complete")
	}
	tt := TruthTable{
		"0 0 1 0 -> 0",
	}
	if tt.Complete() {
		t.Error("tt is supposed to be incomplete")
	}
}

func TestInvert(t *testing.T) {
	pt1 := xorInverted
	pt2 := xor.Invert()
	if len(*pt1) != len(*pt2) {
		t.Error("length of probability tables differ")
	}
	// For each input, check that the output in both pt1 and pt2 matches
	for _, pt1row := range *pt1 {
		if !ValidRow(pt1row) {
			t.Error("xorInverted has an invalid row: " + pt1row)
		}
		output := strings.Split(pt1row, "->")[1]
		if output == "" {
			t.Error("xorInverted has empty output: " + pt1row)
		}
		found := false
		for _, pt2row := range *pt2 {
			if !ValidRow(pt2row) {
				t.Error("xor.Invert() has an invalid row: " + pt2row)
			}
			if strings.Split(pt2row, "->")[1] == output {
				found = true
				break
			}
		}
		if !found {
			t.Error("xorInverted differs from xor.Invert()")
		}
	}
}
