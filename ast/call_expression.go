package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
	// The CallExpression type represents a function call in the abstract syntax
	// tree. For example:
	// add(1, 2)
	CallExpression struct {
		Token     *token.Token
		Function  Node
		Arguments []Node
		Awaited   bool
	}
)

func (ce *CallExpression) String() string {
	var out strings.Builder

	out.WriteString(ce.Function.String())
	out.WriteString("(")

	for i, arg := range ce.Arguments {
		out.WriteString(arg.String())

		if i != len(ce.Arguments)-1 {
			out.WriteString(",")
		}
	}

	out.WriteString(")")
	return out.String()
}
