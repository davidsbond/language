package ast

import (
	"strconv"

	"github.com/davidsbond/dave/token"
)

type (
	// The BooleanLiteral type represents a literal bool within the source code.
	// For example, in a variable assignment:
	// 	var x = true
	BooleanLiteral struct {
		Token *token.Token
		Value bool
	}
)

// TokenLiteral returns the literal value of the token for this
// statement.
func (bl *BooleanLiteral) TokenLiteral() string {
	return bl.Token.Literal
}

func (bl *BooleanLiteral) String() string {
	return strconv.FormatBool(bl.Value)
}
