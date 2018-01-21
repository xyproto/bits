package bits

import (
	"bytes"
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
		digits := strings.Split(strings.TrimSpace(element), " ")
		var b Bits
		for _, digit := range digits {
			if digit == "1" || digit == "0" {
				b = append(b, NewBit(digit))
			} else {
				panic("Invalid choice string: " + s)
			}
		}
		if len(b) > 0 {
			c = append(c, b)
		}
	}
	return &c
}

// Takes a string like this: "0 0 | 0 1 | 1 0 | 1 1" and returns the four elements (separated by "|") as Choices (slice of Bits)
func StringToChoices(s string) Choices {
	c := NewChoices(s)
	return *c
}
