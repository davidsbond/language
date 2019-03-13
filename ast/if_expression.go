package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
	// The IfExpression type represents an if statement in the source code. It consists
	// of a condition, consequence and optional consequence.
	// For example:
	// if (a == 0) { <-- Condition
	//   return -1   <-- Consequence
	// } else {
	//   return a    <-- Alternative
	// }
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
