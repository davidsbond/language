package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	// The ReturnStatement type represents a return statement in the abstract
	// syntax tree. For example
	// return 1 + 1
	ReturnStatement struct {
		Token       *token.Token
		ReturnValue Node
	}
)

func (rs *ReturnStatement) String() string {
	return fmt.Sprintf("return %s", rs.ReturnValue.String())
}
