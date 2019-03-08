package ast

import (
	"fmt"

	"github.com/davidsbond/dave/token"
)

type (
	// The AtomicStatement type represents a atomicant variable assignment
	// in the source code.
	//
	// Example:
	// 	atomic x = 1
	AtomicStatement struct {
		Token *token.Token
		Name  *Identifier
		Value Node
	}
)

// TokenLiteral returns the literal value of the token for this
// statement.
func (vs *AtomicStatement) TokenLiteral() string {
	return vs.Token.Literal
}

func (vs *AtomicStatement) String() string {
	return fmt.Sprintf("atomic %s = %v", vs.Name.Value, vs.Value.String())
}
