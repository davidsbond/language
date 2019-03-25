package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/lexer"
	"github.com/davidsbond/language/parser"
)

type (
	ParserTest struct {
		Name         string
		Expression   string
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

		if _, ok := pt.ExpectedNode.(*ast.HashLiteral); ok {
			for _, ch := range pt.ExpectedNode.String() {
				if !strings.Contains(node.String(), string(ch)) {
					t.Fatalf("expected rune %v in %s, but didn't find it", ch, node.String())
					return
				}
			}
		} else if pt.ExpectedNode.String() != node.String() {
			t.Fatalf("expected %s, got %s", pt.ExpectedNode.String(), node.String())
		}
	})
}
