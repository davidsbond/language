package parser_test

import (
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func TestParser_StringLiteral(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse string literals",
			Expression: `"test"`,
			ExpectedNode: &ast.StringLiteral{
				Token: token.New("1", token.STRING, 0, 0),
				Value: "test",
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
