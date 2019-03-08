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

func TestParser_NumberLiteral(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name            string
		Expression      string
		ExpectedLiteral *ast.NumberLiteral
	}{
		{
			Name:       "It should parse number literals",
			Expression: "1",
			ExpectedLiteral: &ast.NumberLiteral{
				Token: token.New("1", token.NUMBER, 0, 0),
				Value: 1,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)

			result := parser.Parse()

			assert.Len(t, result.Nodes, 1)

			stmt, ok := result.Nodes[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			lit, ok := stmt.Expression.(*ast.NumberLiteral)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedLiteral.String(), lit.String())
		})
	}
}
