package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/token"
	"github.com/stretchr/testify/assert"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
)

func TestParser_ArrayLiteral(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name            string
		Expression      string
		ExpectedLiteral *ast.ArrayLiteral
	}{
		{
			Name:       "It should parse valid array literals",
			Expression: `[1, "test", 't']`,
			ExpectedLiteral: &ast.ArrayLiteral{
				Token: token.New(token.TRUE, token.TRUE, 0, 0),
				Elements: []ast.Node{
					&ast.NumberLiteral{
						Token: token.New("1", token.NUMBER, 0, 0),
						Value: 1,
					},
					&ast.StringLiteral{
						Token: token.New("test", token.STRING, 0, 0),
						Value: "test",
					},
					&ast.CharacterLiteral{
						Token: token.New("t", token.CHAR, 0, 0),
						Value: 't',
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

			stmt, ok := result.Nodes[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			lit, ok := stmt.Expression.(*ast.ArrayLiteral)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedLiteral.String(), lit.String())
		})
	}
}
