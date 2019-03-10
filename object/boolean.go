package object

import (
	"strconv"
)

const (
	// TypeBoolean is the type for a boolean value
	TypeBoolean = "Boolean"
)

type (
	// The Boolean type represents a true/false value in memory.
	Boolean struct {
		Value bool
	}
)

// Type returns the type for this object.
func (b *Boolean) Type() Type {
	return TypeBoolean
}

// Clone creates a copy of the current object that can be used
// without modifying the original value
func (b *Boolean) Clone() Object {
	return &Boolean{Value: b.Value}
}

// HashKey creates a unique identifier for this value for use
// in hashes.
func (b *Boolean) HashKey() HashKey {
	var value float64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (b *Boolean) String() string {
	return strconv.FormatBool(b.Value)
}
