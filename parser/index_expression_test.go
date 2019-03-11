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

func TestParser_IndexExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name               string
		Expression         string
		ExpectedExpression *ast.IndexExpression
	}{
		{
			Name:       "It should parse number index expressions",
			Expression: `test[1]`,
			ExpectedExpression: &ast.IndexExpression{
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
		t.Run(tc.Name, func(t *testing.T) {
			rd := bufio.NewReader(strings.NewReader(tc.Expression))
			lex, _ := lexer.New(rd)
			parser := parser.New(lex)
			result, _ := parser.Parse()

			assert.Len(t, result.Nodes, 1)

			stmt, ok := result.Nodes[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			lit, ok := stmt.Expression.(*ast.IndexExpression)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedExpression.String(), lit.String())
		})
	}
}
