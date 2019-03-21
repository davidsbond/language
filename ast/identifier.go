package ast

import "github.com/davidsbond/language/token"

type (
	// The Identifier type represents an identifier in source code. For example, a variable
	// name or keyword.
	Identifier struct {
		Token *token.Token
		Value string
	}
)

func (i *Identifier) String() string {
	return i.Token.Literal
}
