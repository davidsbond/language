package ast

import "github.com/davidsbond/dave/token"

type (
	// The NumberLiteral type represents a literal number within the source code.
	// For example, in a variable assignment:
	// 	var x = 1
	// The literal value of '1' is stored in the NumberLiteral type.
	NumberLiteral struct {
		Token *token.Token
		Value float64
	}
)

// TokenLiteral returns the literal value of the token for this
// statement.
func (nl *NumberLiteral) TokenLiteral() string {
	return nl.Token.Literal
}

func (nl *NumberLiteral) String() string {
	return nl.Token.Literal
}
