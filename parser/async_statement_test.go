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

func TestParser_AsyncStatement(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name              string
		Expression        string
		ExpectedStatement *ast.AsyncStatement
	}{
		{
			Name: "It should parse valid async statements",
			Expression: `
			async func add(a, b) {
				return a + b
			}
			`,
			ExpectedStatement: &ast.AsyncStatement{
				Token: token.New("return", token.RETURN, 0, 0),
				Value: &ast.FunctionLiteral{
					Name: &ast.Identifier{
						Token: token.New("add", token.IDENT, 0, 0),
						Value: "add",
					},
					Parameters: []*ast.Identifier{
						&ast.Identifier{
							Token: token.New("a", token.IDENT, 0, 0),
							Value: "a",
						},
						&ast.Identifier{
							Token: token.New("b", token.IDENT, 0, 0),
							Value: "b",
						},
					},
					Body: &ast.BlockStatement{
						Token: token.New("{", token.LBRACE, 0, 0),
						Statements: []ast.Node{
							&ast.ReturnStatement{
								Token: token.New(token.RETURN, token.RETURN, 0, 0),
								ReturnValue: &ast.InfixExpression{
									Token:    token.New(token.PLUS, token.PLUS, 0, 0),
									Operator: "+",
									Left: &ast.Identifier{
										Token: token.New("a", token.IDENT, 0, 0),
										Value: "a",
									},
									Right: &ast.Identifier{
										Token: token.New("b", token.IDENT, 0, 0),
										Value: "b",
									},
								},
							},
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

			stmt, ok := result.Nodes[0].(*ast.AsyncStatement)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedStatement.String(), stmt.String())
		})
	}
}
