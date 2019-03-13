package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
	// The IndexExpression type represents an index expression in the source code.
	// For example:
	// var test = "test"
	// var e = test[1]
	IndexExpression struct {
		Token *token.Token
		Left  Node
		Index Node
	}
)

func (ie *IndexExpression) String() string {
	var out strings.Builder

	out.WriteString(ie.Left.String())
	out.WriteRune('[')
	out.WriteString(ie.Index.String())
	out.WriteRune(']')

	return out.String()
}
