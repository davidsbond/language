package ast

import (
	"strings"

	"github.com/davidsbond/language/token"
)

type (
	// The BlockStatement type represents statements wrapped in braces.
	BlockStatement struct {
		Token      *token.Token
		Statements []Node
	}
)

func (bl *BlockStatement) String() string {
	var out strings.Builder

	for _, stmt := range bl.Statements {
		out.WriteString("\t")
		out.WriteString(stmt.String())
	}

	return out.String()
}
