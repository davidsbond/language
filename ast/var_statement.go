package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	// The VarStatement type represents a mutable variable assignment
	// in the source code.
	//
	// Example:
	// 	var x = 1
	VarStatement struct {
		Token *token.Token
		Name  *Identifier
		Value Node
	}
)

// TokenLiteral returns the literal value of the token for this
// statement.
func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}

func (vs *VarStatement) String() string {
	return fmt.Sprintf("var %s = %v", vs.Name.Value, vs.Value.String())
}
