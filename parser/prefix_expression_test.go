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

func TestParser_PrefixExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name               string
		Expression         string
		ExpectedExpression *ast.PrefixExpression
	}{
		{
			Name:       "It should parse negative prefixes",
			Expression: "-1",
			ExpectedExpression: &ast.PrefixExpression{
				Token:    token.New(token.MINUS, token.MINUS, 0, 0),
				Operator: "-",
				Right: &ast.NumberLiteral{
					Token: token.New("1", token.NUMBER, 0, 0),
					Value: 1,
				},
			},
		},
		{
			Name:       "It should parse bang prefixes",
			Expression: "!true",
			ExpectedExpression: &ast.PrefixExpression{
				Token:    token.New(token.BANG, token.BANG, 0, 0),
				Operator: "!",
				Right: &ast.BooleanLiteral{
					Token: token.New("true", token.TRUE, 0, 0),
					Value: true,
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

			exp, ok := stmt.Expression.(*ast.PrefixExpression)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedExpression.String(), exp.String())
		})
	}
}
