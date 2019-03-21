package ast

import (
	"strings"

	"github.com/davidsbond/language/token"
)

type (
	// The ArrayLiteral type represents a literal array object in source code.
	// For example:
	// const arr = [1, 2, 3, 4]
	ArrayLiteral struct {
		Token    *token.Token
		Elements []Node
	}
)

func (al *ArrayLiteral) String() string {
	var out strings.Builder

	out.WriteByte('[')

	for i, elem := range al.Elements {
		out.WriteString(elem.String())

		if i != len(al.Elements)-1 {
			out.WriteString(", ")
		}
	}

	out.WriteByte(']')

	return out.String()
}
