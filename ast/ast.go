// Package ast contains types that are contained within the abstract
// syntax tree for the language
package ast

import (
	"strings"
)

const (
	// NodeTypeStatement denotes a statement node
	NodeTypeStatement = iota
	// NodeTypeExpression denotes an expression node
	NodeTypeExpression
	// NodeTypeAST denotes the root node of an AST
	NodeTypeAST
)

type (
	// The AST type contains the tree representation of the source code.
	AST struct {
		Nodes []Node
	}

	// The Node interface defines methods used by types that can be found in
	// the abstract syntax tree.
	Node interface {
		TokenLiteral() string
		String() string
	}
)

func (ast *AST) String() string {
	var builder strings.Builder

	for _, node := range ast.Nodes {
		builder.WriteString(node.String())
		builder.WriteByte('\n')
	}

	return builder.String()
}

// TokenLiteral returns blank for the AST. It's just to satisfy the interface.
func (ast *AST) TokenLiteral() string {
	return ""
}
