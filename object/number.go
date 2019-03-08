package object

import (
	"fmt"
)

const (
	// TypeNumber is the type for a numerical value
	TypeNumber = "Number"
)

type (
	// The Number type represents a number stored in memory. All numbers are 64 bit
	// floating point numbers.
	Number struct {
		Value float64
	}
)

// Type returns the type of the object.
func (num *Number) Type() Type {
	return TypeNumber
}

// Clone creates a copy of the current object that can be used
// without modifying the original value
func (num *Number) Clone() Object {
	return &Number{Value: num.Value}
}

func (num *Number) String() string {
	return fmt.Sprint(num.Value)
}
