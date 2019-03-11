package ast

import (
	"strings"

	"github.com/davidsbond/dave/token"
)

type (
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
