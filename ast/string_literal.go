package ast

import (
	"fmt"

	"github.com/davidsbond/language/token"
)

type (
	// The StringLiteral type represents a literal string within the source code.
	// For example, in a variable assignment:
	// 	var x = "test"
	// The literal value of 'test' is stored in the StringLiteral type.
	StringLiteral struct {
		Token *token.Token
		Value string
	}
)

func (nl *StringLiteral) String() string {
	return fmt.Sprintf(`"%s"`, nl.Value)
}
