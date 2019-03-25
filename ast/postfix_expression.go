package ast

import (
	"fmt"

	"github.com/davidsbond/language/token"
)

type (
	// The PostfixExpression type represents a postfix expression in the source
	// code. For example:
	// var a = 1
	// a++
	// a--
	PostfixExpression struct {
		Token    *token.Token
		Left     *Identifier
		Operator string
	}
)

func (pe *PostfixExpression) String() string {
	return fmt.Sprintf("%s%s", pe.Left.String(), pe.Operator)
}
