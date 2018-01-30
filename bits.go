package bits

import (
	"bytes"
	"errors"
	"strings"
)

// Shortcuts for bits. Remember to compare a Bit with a Bit, not with an int!
const (
	B0 = Bit(0)
	B1 = Bit(1)
)

// Take a slice of ints (0 or 1), return a space separated string ("0 1 0 1 0")
func (b Bits) String() string {
	var buf bytes.Buffer
	lastPos := len(b) - 1
	for i := 0; i <= lastPos; i++ {
		buf.WriteString(b[i].String())
		if i < lastPos {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}

// Convert a single bit to a string
func (b Bit) String() string {
	if b == B1 {
		return "1"
	}
	return "0"
}

// Convert a "0" or "1" to a single bit
func NewBit(s string) Bit {
	if s == "1" {
		return B1
	}
	return B0
}

// Check if two slices of Bit are equal
func (b *Bits) Equal(x *Bits) bool {
	bLen := len(*b)
	xLen := len(*x)
	if bLen != xLen {
		return false
	}
	for i := 0; i < bLen; i++ {
		if (*b)[i] != (*x)[i] {
			return false
		}
	}
	return true
}

// StringToBits converts a space-separated string of "0"s and "1"s to a pointer
// to a slice of bits.Bit. Returns an error if an invalid string is given.
func StringToBits(s string) (*Bits, error) {
	var gatheredBits Bits
	if strings.Contains(s, " ") {
		for _, sbit := range strings.Split(s, " ") {
			if sbit == "1" {
				gatheredBits = append(gatheredBits, B1)
			} else if sbit == "0" {
				gatheredBits = append(gatheredBits, B0)
			} else {
				return nil, errors.New("Invalid bit: " + sbit)
			}
		}
	} else {
		if s == "1" {
			gatheredBits = append(gatheredBits, B1)
		} else if s == "0" {
			gatheredBits = append(gatheredBits, B0)
		} else {
			return nil, errors.New("Invalid bit: " + s)
		}
	}
	return &gatheredBits, nil
}

// Takes a string like this: "0 0 | 0 1 | 1 0 | 1 1" and returns the four elements (separated by "|") as *Choices (slice of Bits)
func NewChoices(s string) *Choices {
	if strings.Contains(s, "->") {
		panic("choice string should not contain ->: " + s)
	}
	var c Choices
	for _, element := range strings.Split(s, "|") {
		if !strings.Contains(element, " ") {
			continue
		}
		b, err := StringToBits(strings.TrimSpace(element))
		if err != nil {
			panic("Invalid choice string: " + s)
		}
		if len(*b) > 0 {
			c = append(c, *b)
		}
	}
	return &c
}

// Takes a string like this: "0 0 | 0 1 | 1 0 | 1 1" and returns the four elements (separated by "|") as Choices (slice of Bits)
func StringToChoices(s string) Choices {
	c := NewChoices(s)
	return *c
}
