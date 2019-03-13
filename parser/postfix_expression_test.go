package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_PostfixExpression(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse incremental postfixes",
			Expression: "a++",
			ExpectedNode: &ast.PostfixExpression{
				Token:    token.New(token.INC, token.INC, 0, 0),
				Operator: "++",
				Left: &ast.Identifier{
					Token: token.New("a", token.IDENT, 0, 0),
					Value: "a",
				},
			},
		},
		{
			Name:       "It should parse decremental postfixes",
			Expression: "a--",
			ExpectedNode: &ast.PostfixExpression{
				Token:    token.New(token.DEC, token.DEC, 0, 0),
				Operator: "--",
				Left: &ast.Identifier{
					Token: token.New("a", token.IDENT, 0, 0),
					Value: "a",
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
