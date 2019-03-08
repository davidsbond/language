package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
	"github.com/davidsbond/dave/token"
	"github.com/stretchr/testify/assert"
)

func TestParser_FunctionLiteral(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name            string
		Expression      string
		ExpectedLiteral *ast.FunctionLiteral
	}{
		{
			Name: "It should parse a valid function literal",
			Expression: `
			func add(a, b) {
				return a + b
			}`,
			ExpectedLiteral: &ast.FunctionLiteral{
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
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)

			result, _ := parser.Parse()

			assert.Len(t, result.Nodes, 1)

			stmt, ok := result.Nodes[0].(*ast.FunctionLiteral)

			assert.True(t, ok)
			assert.Equal(t, tc.ExpectedLiteral.String(), stmt.String())
		})
	}
}
