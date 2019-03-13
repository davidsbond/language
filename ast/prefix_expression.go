package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	// The PrefixExpression type represents a prefix expression within
	// the source code. For example:
	// var a = true
	// var b = !a
	PrefixExpression struct {
		Token    *token.Token
		Operator string
		Right    Node
	}
)

func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("%s%s", pe.Operator, pe.Right.String())
}
