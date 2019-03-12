package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
)

type (
	ParserTest struct {
		Name               string
		Expression         string
		ExpectedNode ast.Node
	}
)

func (pt *ParserTest) Run(t *testing.T) {
	t.Run(pt.Name, func(t *testing.T) {
		rd := bufio.NewReader(strings.NewReader(pt.Expression))
		lex, _ := lexer.New(rd)
		parser := parser.New(lex)

		result, _ := parser.Parse()
		node := result.Nodes[0]

		if pt.ExpectedNode.String() != node.String() {
			t.Fatalf("expected %s, got %s", pt.ExpectedNode.String(), node.String())
		}
	})
}
