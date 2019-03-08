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

func (bl *BooleanLiteral) String() string {
	return strconv.FormatBool(bl.Value)
}
