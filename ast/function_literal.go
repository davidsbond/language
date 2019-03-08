package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
	FunctionLiteral struct {
		Token      *token.Token
		Name       *Identifier
		Parameters []*Identifier
		Body       *BlockStatement
	}
)

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out strings.Builder

	out.WriteString("function ")
	out.WriteString(fl.Name.String())
	out.WriteString("(")

	for i, ident := range fl.Parameters {
		out.WriteString(ident.String())

		if i != len(fl.Parameters)-1 {
			out.WriteString(",")
		}
	}

	out.WriteString(") {\n")
	out.WriteString(fl.Body.String())
	out.WriteString("\n}")

	return out.String()
}
