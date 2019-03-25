package parser_test

import (
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func TestParser_BooleanLiteral(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse true literals",
			Expression: "true",
			ExpectedNode: &ast.BooleanLiteral{
				Token: token.New(token.TRUE, token.TRUE, 0, 0),
				Value: true,
			},
		},
		{
			Name:       "It should parse false literals",
			Expression: "false",
			ExpectedNode: &ast.BooleanLiteral{
				Token: token.New(token.FALSE, token.FALSE, 0, 0),
				Value: false,
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
