package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_PrefixExpression(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse negative prefixes",
			Expression: "-1",
			ExpectedNode: &ast.PrefixExpression{
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Operator: "-",
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse bang prefixes",
			Expression: "!true",
			ExpectedNode: &ast.PrefixExpression{
				Token:    token.New(token.BANG, token.BANG, 0, 0),
				Operator: "!",
				Right: &ast.BooleanLiteral{
					Token: token.New("true", token.TRUE, 0, 0),
					Value: true,
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
