package parser_test

import (
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func TestParser_IndexExpression(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse number index expressions",
			Expression: `test[1]`,
			ExpectedNode: &ast.IndexExpression{
				Token: token.New("[", token.LBRACKET, 0, 0),
				Left: &ast.Identifier{
					Token: token.New("test", token.IDENT, 0, 0),
					Value: "test",
				},
				Index: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
