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

func TestParser_BooleanLiteral(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name            string
		Expression      string
		ExpectedLiteral *ast.BooleanLiteral
	}{
		{
			Name:       "It should parse true literals",
			Expression: "true",
			ExpectedLiteral: &ast.BooleanLiteral{
				Token: token.New(token.TRUE, token.TRUE, 0, 0),
				Value: true,
			},
		},
		{
			Name:       "It should parse false literals",
			Expression: "false",
			ExpectedLiteral: &ast.BooleanLiteral{
				Token: token.New(token.FALSE, token.FALSE, 0, 0),
				Value: false,
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

			lit, ok := stmt.Expression.(*ast.BooleanLiteral)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedLiteral.String(), lit.String())
		})
	}
}
