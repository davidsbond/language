package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_NumberLiteral(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse number literals",
			Expression: "1",
			ExpectedNode: &ast.NumberLiteral{
				Token: token.New("1", token.NUMBER, 0, 0),
				Value: 1,
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
