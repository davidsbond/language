package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
	// The FunctionLiteral type represents a function declaration in the abstract
	// syntax tree. For example:
	// function add(a, b) { return a + b }
	FunctionLiteral struct {
		Token      *token.Token
		Name       *Identifier
		Parameters []*Identifier
		Body       *BlockStatement
	}
)

func (fl *FunctionLiteral) String() string {
	var out strings.Builder

	if fl.Name.Value == "anonymous" {
		out.WriteString("func")
	} else {
		out.WriteString("func ")
		out.WriteString(fl.Name.String())
	}

	out.WriteString("(")

	for i, ident := range fl.Parameters {
		out.WriteString(ident.String())

		if i != len(fl.Parameters)-1 {
			out.WriteString(",")
		}
	}

	out.WriteString(") {\n")
	out.WriteString(fl.Body.String())
	out.WriteString("\n}\n")

	return out.String()
}
