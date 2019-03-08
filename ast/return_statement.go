package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	ReturnStatement struct {
		Token       *token.Token
		ReturnValue Node
	}
)

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	return fmt.Sprintf("return %s", rs.ReturnValue.String())
}
