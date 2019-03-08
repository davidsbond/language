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

func (vs *ConstStatement) String() string {
	return fmt.Sprintf("const %s = %v", vs.Name.Value, vs.Value.String())
}
