package ast

import "github.com/davidsbond/dave/token"

type (
	// The Identifier type represents an identifier in source code. For example, a variable
	// name or keyword.
	Identifier struct {
		Token *token.Token
		Value string
	}
)

// TokenLiteral returns the literal value of the token for this
// statement.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Token.Literal
}
