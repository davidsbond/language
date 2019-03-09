package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
	HashLiteral struct {
		Token *token.Token
		Pairs map[Node]Node
	}
)

func (hl *HashLiteral) String() string {
	var out strings.Builder

	out.WriteString("{\n")

	for key, val := range hl.Pairs {
		out.WriteRune('\t')
		out.WriteString(key.String())
		out.WriteString(": ")
		out.WriteString(val.String())
		out.WriteRune('\n')
	}

	out.WriteString("}\n")

	return out.String()
}
