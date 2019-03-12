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

func TestParser_PostfixExpression(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name               string
		Expression         string
		ExpectedExpression *ast.PostfixExpression
	}{
		{
			Name:       "It should parse incremental postfixes",
			Expression: "a++",
			ExpectedExpression: &ast.PostfixExpression{
				Token:    token.New(token.INC, token.INC, 0, 0),
				Operator: "++",
				Left: &ast.Identifier{
					Token: token.New("a", token.IDENT, 0, 0),
					Value: "a",
				},
			},
		},
		{
			Name:       "It should parse decremental postfixes",
			Expression: "a--",
			ExpectedExpression: &ast.PostfixExpression{
				Token:    token.New(token.DEC, token.DEC, 0, 0),
				Operator: "--",
				Left: &ast.Identifier{
					Token: token.New("a", token.IDENT, 0, 0),
					Value: "a",
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

			exp, ok := stmt.Expression.(*ast.PostfixExpression)
			assert.True(t, ok)

			assert.Equal(t, tc.ExpectedExpression.String(), exp.String())
		})
	}
}
