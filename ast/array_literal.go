package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
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
