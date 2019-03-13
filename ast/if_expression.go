package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
	IfExpression struct {
		Token       *token.Token
		Condition   Node
		Consequence *BlockStatement
		Alternative *BlockStatement
	}
)

func (ie *IfExpression) String() string {
	var out strings.Builder

	out.WriteString("if (")
	out.WriteString(ie.Condition.String())
	out.WriteString(") {\n")
	out.WriteString(ie.Consequence.String())
	out.WriteString("\n}\n")

	if ie.Alternative != nil {
		out.WriteString("else {\n")
		out.WriteString(ie.Alternative.String())
		out.WriteString("\n}\n")
	}

	return out.String()
}
