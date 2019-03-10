package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	PrefixExpression struct {
		Token    *token.Token
		Operator string
		Right    Node
	}
)

func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("%s%s", pe.Operator, pe.Right.String())
}
