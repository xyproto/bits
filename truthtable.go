package bits

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type TruthTable []string
type ProbTable []string

type ValidatorFunc func(*Bits) bool

// sort the words in a given string and return a new string
func sSort(s string) string {
	l := strings.Split(s, " ")
	sort.Strings(l)
	return strings.Join(l, " ")
}

// Make it possible to run a truth table in reverse,
// by returning a table with probabilities
func (tt *TruthTable) Invert() *ProbTable {
	iomap := make(map[string]string)
	for _, row := range *tt {
		fields := strings.Split(row, "->")
		if len(fields) != 2 {
			panic("Invalid row: " + row)
		}
		inputs := strings.TrimSpace(fields[0])
		outputs := strings.TrimSpace(fields[1])
		if _, hasKey := iomap[outputs]; hasKey {
			iomap[outputs] += " | " + inputs
		} else {
			iomap[outputs] = inputs
		}
	}
	var pt ProbTable
	for outputs, inputs := range iomap {
		pt = append(pt, outputs+" -> "+inputs)
	}
	return &pt
}

// Check if a row in a truth table has valid syntax
func ValidRow(ttRow string) bool {
	// No double space
	if strings.Contains(ttRow, "  ") {
		return false
	}
	// Must have an arrow
	if strings.Count(ttRow, "->") != 1 {
		return false
	}
	// Only 0, 1, space, |, - and > are allowed characters
	for _, r := range ttRow {
		switch r {
		case '0', '1', ' ', '|', '-', '>':
		default:
			// Invalid rune
			return false
		}
	}
	// Acceptable
	return true
}

// Check if a truth table is complete
// If tables are assumed to not contain duplicates, one could just count the rows too
func (tt *TruthTable) Complete() bool {
	if len(*tt) == 0 {
		return false
	}
	ttRow := (*tt)[0]
	if !ValidRow(ttRow) {
		panic("Not a valid row: " + ttRow)
	}
	inputLen := len(strings.Split(strings.TrimSpace(strings.Split(ttRow, "->")[0]), " "))
	// If the input length is 3, the numbers to cycle over are 0..2**3-1
	for i := 0; float64(i) < math.Pow(2.0, float64(inputLen)); i++ {
		bitSlice := strings.Split(fmt.Sprintf("%b", i), "")
		// Padding the slice with zeros, on the left side
		for len(bitSlice) < inputLen {
			bitSlice = append([]string{"0"}, bitSlice...)
		}
		bits := strings.Join(bitSlice, " ")
		// OK, have generated a string of bits, like: "0 1 0 0 1 1 0"
		found := false
		for _, ttRow = range *tt {
			if strings.Contains(ttRow, bits) {
				found = true
				break
			}
		}
		if !found {
			//fmt.Println("Could not find: " + bits + " in truth table")
			return false
		}
	}
	return true
}

// Take truth table, return a Gate
func (tt *TruthTable) Gate() OneToManyGate {
	return func(inputs Bits) Bit {
		inputRow := inputs.String()
		for _, ttRow := range *tt {
			if !ValidRow(ttRow) {
				panic("Not a valid row: " + ttRow)
			}
			if strings.Contains(ttRow, inputRow) {
				// Found the inputs in the table, can now return the output
				v, err := strconv.Atoi(strings.TrimSpace(strings.Split(ttRow, "->")[1]))
				if err != nil {
					panic("Rows in a truth table must always end with an output bit. Current: " + ttRow)
				}
				// Return the corresponding output bit
				return Bit(v)
			}
		}
		// The truth table is incomplete, panic
		panic("Truth table is incomplete, no row for: " + inputRow)
	}
}

func (tt *TruthTable) String() string {
	return strings.Join(*tt, ", ")
}

// Take prob table, return a prob gate
func (pt *ProbTable) Gate() ProbGate {
	return func(input Bit) Choices {
		inputRow := input.String()
		for _, ttRow := range *pt {
			if !ValidRow(ttRow) {
				panic("Not a valid row: " + ttRow)
			}
			if strings.HasPrefix(strings.TrimSpace(ttRow), strings.TrimSpace(inputRow)) {
				// Found the inputs in the table, can now return the output
				return StringToChoices(strings.TrimSpace(strings.Split(ttRow, "->")[1]))
			}
		}
		// The truth table is incomplete, panic
		panic("Truth table is incomplete, no row for: " + inputRow)
	}
}

func (pt *ProbTable) String() string {
	return strings.Join(*pt, ", ")
}

// Given a truth table and an output bit, reverse the output and find the input.
// Also needs a ValidatorFunc to return true when the correct input has been found.
// Returns the found input, number of iterations that were required and an error if not found
func Reverse(tt *TruthTable, output Bit, vf ValidatorFunc) (foundInput *Bits, iterations uint, err error) {
	var counter uint
	for _, possibleInput := range tt.Invert().Process(output) {
		counter++
		if vf(&possibleInput) {
			return &possibleInput, counter, nil
		}
	}
	return nil, counter, errors.New("Unable to reverse " + tt.String())
}
