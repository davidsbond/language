package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_AsyncStatement(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name: "It should parse valid async statements",
			Expression: `
			async func add(a, b) {
				return a + b
			}
			`,
			ExpectedNode: &ast.AsyncStatement{
				Token: token.New("return", token.RETURN, 0, 0),
				Value: &ast.FunctionLiteral{
					Name: &ast.Identifier{
						Token: token.New("add", token.IDENT, 0, 0),
						Value: "add",
					},
					Parameters: []*ast.Identifier{
						{
							Token: token.New("a", token.IDENT, 0, 0),
							Value: "a",
						},
						{
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
		tc.Run(t)
	}
}
