package object

import "strings"

const (
	// TypeArray is the type for an array value
	TypeArray = "Array"
)

type (
	// The Array type represents an array of objects
	Array struct {
		Elements []Object
	}
)

// Type returns this object's type.
func (ar *Array) Type() Type {
	return TypeArray
}

// Clone returns a copy of the array as a new pointer.
func (ar *Array) Clone() Object {
	return &Array{Elements: ar.Elements}
}

func (ar *Array) String() string {
	var out strings.Builder

	out.WriteByte('[')

	for i, elem := range ar.Elements {
		out.WriteString(elem.String())

		if i != len(ar.Elements)-1 {
			out.WriteString(", ")
		}
	}

	out.WriteByte(']')

	return out.String()
}
