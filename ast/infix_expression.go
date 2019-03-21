package ast

import (
	"bytes"

	"github.com/davidsbond/language/token"
)

type (
	// The InfixExpression type represents an infix expression in the source code. This is where
	// an operator exists between two expressions, such as 1 + 1 or 2 * 2.
	InfixExpression struct {
		Token    *token.Token
		Left     Node
		Operator string
		Right    Node
	}
)

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())

	return out.String()
}
