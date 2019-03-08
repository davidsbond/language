package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/lexer"
	"github.com/davidsbond/dave/parser"
	"github.com/davidsbond/dave/token"
	"github.com/stretchr/testify/assert"
)

func TestParser_CallExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name               string
		Expression         string
		ExpectedExpression *ast.CallExpression
	}{
		{
			Name:       "It should parse a valid call expression",
			Expression: "test(1, 2, 3)",
			ExpectedExpression: &ast.CallExpression{
				Function: &ast.Identifier{
					Value: "test",
					Token: token.New("test", token.IDENT, 0, 0),
				},
				Arguments: []ast.Node{
					&ast.NumberLiteral{
						Token: token.New("1", token.NUMBER, 0, 0),
						Value: 1,
					},
					&ast.NumberLiteral{
						Token: token.New("2", token.NUMBER, 0, 0),
						Value: 2,
					},
					&ast.NumberLiteral{
						Token: token.New("3", token.NUMBER, 0, 0),
						Value: 3,
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

			exp, ok := stmt.Expression.(*ast.CallExpression)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedExpression.String(), exp.String())
		})
	}
}
