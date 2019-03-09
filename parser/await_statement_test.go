package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/token"
	"github.com/stretchr/testify/assert"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
)

func TestParser_AwaitStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name              string
		Expression        string
		ExpectedStatement *ast.AwaitStatement
	}{
		{
			Name: "It should parse valid await statements",
			Expression: `
			await add(1, 2)
			`,
			ExpectedStatement: &ast.AwaitStatement{
				Token: token.New("await", token.AWAIT, 0, 0),
				Value: &ast.CallExpression{
					Function: &ast.Identifier{
						Value: "add",
						Token: token.New("add", token.IDENT, 0, 0),
					},
					Arguments: []ast.Node{
						&ast.NumberLiteral{
							Token: token.New("1", token.NUMBER, 0, 0),
							Value: 1,
						},
						&ast.NumberLiteral{
							Token: token.New("2", token.NUMBER, 0, 0),
							Value: 2,
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)

			result, _ := parser.Parse()

			assert.Len(t, result.Nodes, 1)

			call, ok := result.Nodes[0].(*ast.AwaitStatement)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedStatement.String(), call.String())
		})
	}
}
