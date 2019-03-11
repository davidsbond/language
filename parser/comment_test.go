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

func TestParser_Comment(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name            string
		Expression      string
		ExpectedComment *ast.Comment
	}{
		{
			Name:       "It should parse comments",
			Expression: "// A test comment",
			ExpectedComment: &ast.Comment{
				Token: token.New(" A test comment", token.COMMENT, 0, 0),
				Value: "A test comment",
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

			cmt, ok := result.Nodes[0].(*ast.Comment)

			assert.True(t, ok)
			assert.Equal(t, tc.ExpectedComment.String(), cmt.String())
		})
	}
}
