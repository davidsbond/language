package ast

import "github.com/davidsbond/dave/token"

type (
	// The CharacterLiteral type represents a literal character within the source code.
	// For example, in a variable assignment:
	// 	var x = 't'
	// The literal value of 't' is stored in the CharacterLiteral type.
	CharacterLiteral struct {
		Token *token.Token
		Value rune
	}
)

// TokenLiteral returns the literal value of the token for this
// statement.
func (cl *CharacterLiteral) TokenLiteral() string {
	return cl.Token.Literal
}

func (cl *CharacterLiteral) String() string {
	return string(cl.Value)
}