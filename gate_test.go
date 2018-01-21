package bits

import (
	"fmt"
	"testing"
)

func TestXorGate(t *testing.T) {
	xorGate := xor.Gate()

	if xorGate(Bits{0, 0}) != 0 {
		t.Error("0 xor 0 should be 0")
	}
	if xorGate(Bits{0, 1}) != 1 {
		t.Error("0 xor 1 should be 1")
	}
	if xorGate(Bits{1, 0}) != 1 {
		t.Error("1 xor 0 should be 1")
	}
	if xorGate(Bits{1, 1}) != 0 {
		t.Error("1 xor 1 should be 0")
	}
}

func TestProcess(t *testing.T) {
	inputs := Bits{0, 1}
	output := xor.Process(inputs)

	if output != B1 {
		t.Error("0 xor 1 should be 1, but is: " + fmt.Sprintf("%v (%T)", output, output))
	}
}

func TestProb(t *testing.T) {

	inputs := Bits{1, 1}

	correctOutputBit := xor.Process(inputs)
	possibleInputBitsCollection := xorInverted.Process(correctOutputBit)

	// Try to find the input bits
	for i, possibleInputBits := range possibleInputBitsCollection {
		if inputs.Equal(&possibleInputBits) {
			//fmt.Println(possibleInputBits, "(possible answer) ==", inputs, "(correct answer)")
			//fmt.Printf("Found the answer to xor inverted in %d iterations\n", i)
			if i != 1 {
				t.Error("Wrong number of iterations for finding the original input for input Bits{1, 1}")
			}
			return
			//	} else {
			//		fmt.Println(possibleInputBits, "(possible answer) !=", inputs, "(correct answer)")
		}
	}

	t.Error("Could not recover the original input bits")
}
