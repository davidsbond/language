package parser_test

import (
	"testing"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func TestParser_CharacterLiteral(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse character literals",
			Expression: `'t'`,
			ExpectedNode: &ast.CharacterLiteral{
				Token: token.New("t", token.CHAR, 0, 0),
				Value: 't',
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
