package bits

import (
	"fmt"
)

// A new Gate is defined by a truth table, or by a choicetable
type (
	Bit  int   // bit, as integer
	Bits []Bit // input bits, or output bits

	// Multiple possible outputs, with equal probability
	Choices []Bits

	// A logic gate with multiple inputs and one output
	OneToManyGate func(Bits) Bit

	// A logic gate (like "not")
	OneToOneGate func(Bit) Bit

	// A ManyToManyGate can have several inputs and outputs
	ManyToManyGate func(Bits) Bits

	// A simple probgate with one input and several possible outputs
	// Example prob table: {"0 -> 1 1 | 0 0"}, would return Choices{Outputs{1, 1}, Outputs{0, 0}} if given 0
	ProbGate func(Bit) Choices

	// A probgate with many inputs and several possible outputs
	// Example prob table: {"0 1 1 0 1 -> 1 1 | 0 0"}, would return Choices{Outputs{1, 1}, Outputs{0, 0}} if given 0
	MultiProbGate func(Bits) Choices
)

// A gate must have a Process function, that can take Input or Inputs, and return Output or Outputs
type Gater interface {
	Process(interface{}) interface{}
}

// OneToManyGate can act as a gate (Gater)
// many to one
func (lg *OneToManyGate) Process(inputOrInputs interface{}) (outputOrOutputs interface{}) {
	switch v := inputOrInputs.(type) {
	case Bit:
		return (*lg)(Bits{v})
	case Bits:
		return (*lg)(v)
	default:
		panic(fmt.Sprintf("Invalid input, value: %v, type: %T", inputOrInputs, inputOrInputs))
	}
}

// OneToOneGate can act as a gate (Gater)
// one to one
func (ug *OneToOneGate) Process(inputOrInputs interface{}) (outputOrOutputs interface{}) {
	switch v := inputOrInputs.(type) {
	case Bit:
		return (*ug)(v)
	case Bits:
		if len(v) == 1 {
			return (*ug)(v[0])
		}
		panic("A OneToOneGate can only take one input")
	default:
		panic(fmt.Sprintf("Invalid input, value: %v, type: %T", inputOrInputs, inputOrInputs))
	}
}

// Multi can act as a gate (Gater)
// many to many
func (mg *ManyToManyGate) Process(inputOrInputs interface{}) (outputOrOutputs interface{}) {
	switch v := inputOrInputs.(type) {
	case Bit:
		return (*mg)(Bits{v})
	case Bits:
		return (*mg)(v)
	default:
		panic(fmt.Sprintf("Invalid input, value: %v, type: %T", inputOrInputs, inputOrInputs))
	}
}

// Choice Gate can act as a gate (Gater)
// one bit to many possible choices
func (scg *ProbGate) Process(inputOrInputs interface{}) (outputOrOutputs interface{}) {
	switch v := inputOrInputs.(type) {
	case Bit:
		return (*scg)(v)
	case Bits:
		if len(v) == 1 {
			return (*scg)(v[0])
		}
		panic("A ProbGate can only take one input")
	default:
		panic(fmt.Sprintf("Invalid input, value: %v, type: %T", inputOrInputs, inputOrInputs))
	}
}

// Multi Choice Gate can act as a gate (Gater)
// many bits to many possible choices
func (cg *MultiProbGate) Process(inputOrInputs interface{}) (outputOrOutputs interface{}) {
	switch v := inputOrInputs.(type) {
	case Bit:
		return (*cg)(Bits{v})
	case Bits:
		return (*cg)(v)
	default:
		panic(fmt.Sprintf("Invalid input, value: %v, type: %T", inputOrInputs, inputOrInputs))
	}
}

// Create a new logic gate, like "and" or "xor", with inputs and one output
func NewOneToManyGate(tt *TruthTable) *OneToManyGate {
	g := tt.Gate()
	return &g
}

// Create a new probability gate, a ProbGate, where one input can give many alternative outputs
func NewProbGate(pt *ProbTable) *ProbGate {
	g := pt.Gate()
	return &g
}

// TruthTable can act as a gate (Gater), many to one
func (tt *TruthTable) Process(inputs Bits) Bit {
	lg := NewOneToManyGate(tt)
	if output, ok := lg.Process(inputs).(Bit); ok {
		return output
	}
	panic("Invalid return value from TruthTable Process")
}

// ProbTable can act as a gate (Gater), one to many different choices
func (pt *ProbTable) Process(input Bit) Choices {
	pg := NewProbGate(pt)
	if output, ok := pg.Process(input).(Choices); ok {
		return output
	}
	panic("Invalid return value from ProbTable Process")
}
