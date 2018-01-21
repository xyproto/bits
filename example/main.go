package main

import (
	"fmt"
	"github.com/xyproto/bits"
)

func main() {
	// Define the XOR function with a TruthTable
	xor := &bits.TruthTable{
		"0 0 -> 0",
		"0 1 -> 1",
		"1 0 -> 1",
		"1 1 -> 0",
	}

	// This program reverses XOR from 0 to 1, 1 without trying all 4 possible inputs
	to := bits.Bits{1, 1} // Could be 0 0, 0 1, 1 0 or 1 1
	from := xor.Process(to) // The output of the xor function given 1 1 is 0

	// Create a validator function
	validator := func (input *bits.Bits) bool {
		return input.Equal(&to)
	}

	// Reverse the xor function and find the input bits, as confirmed by the validator
	foundInput, iterations, err := bits.Reverse(xor, from, validator)
	if err != nil {
		panic(err)
	}

	// Output info
	fmt.Printf("Reversed xor in %d iterations (instead of %d) and found %s given %s\n", iterations, len(*xor), foundInput.String(), from.String())
}
