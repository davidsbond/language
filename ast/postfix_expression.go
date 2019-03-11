package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	PostfixExpression struct {
		Token    *token.Token
		Left     *Identifier
		Operator string
	}
)

func (pe *PostfixExpression) String() string {
	return fmt.Sprintf("%s%s", pe.Left.String(), pe.Operator)
}
