package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
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
