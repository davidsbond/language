package object

import (
	"strconv"
)

const (
	// TypeBoolean is the node type for a boolean value
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

func (b *Boolean) String() string {
	return strconv.FormatBool(b.Value)
}
