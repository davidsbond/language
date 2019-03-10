package object

import "strings"

const (
	TypeArray = "Array"
)

type (
	Array struct {
		Elements []Object
	}
)

func (ar *Array) Type() Type {
	return TypeArray
}

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
