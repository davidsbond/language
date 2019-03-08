package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	// The ConstStatement type represents a constant variable assignment
	// in the source code.
	//
	// Example:
	// 	const x = 1
	ConstStatement struct {
		Token *token.Token
		Name  *Identifier
		Value Node
	}
)

// TokenLiteral returns the literal value of the token for this
// statement.
func (vs *ConstStatement) TokenLiteral() string {
	return vs.Token.Literal
}

func (vs *ConstStatement) String() string {
	return fmt.Sprintf("const %s = %v", vs.Name.Value, vs.Value.String())
}
