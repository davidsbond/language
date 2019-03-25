package parser_test

import (
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func TestParser_IfExpression(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse if expressions",
			Expression: "if (true) { return false } else { return true }",
			ExpectedNode: &ast.IfExpression{
				Token: token.New(token.IF, token.IF, 0, 0),
				Condition: &ast.BooleanLiteral{
					Token: token.New(token.TRUE, token.TRUE, 0, 0),
					Value: true,
				},
				Consequence: &ast.BlockStatement{
					Token: token.New(token.LBRACE, token.LBRACE, 0, 0),
					Statements: []ast.Node{
						&ast.ReturnStatement{
							Token: token.New(token.RETURN, token.RETURN, 0, 0),
							ReturnValue: &ast.BooleanLiteral{
								Token: token.New(token.FALSE, token.FALSE, 0, 0),
								Value: false,
							},
						},
					},
				},
				Alternative: &ast.BlockStatement{
					Token: token.New(token.LBRACE, token.LBRACE, 0, 0),
					Statements: []ast.Node{
						&ast.ReturnStatement{
							Token: token.New(token.RETURN, token.RETURN, 0, 0),
							ReturnValue: &ast.BooleanLiteral{
								Token: token.New(token.TRUE, token.TRUE, 0, 0),
								Value: true,
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
