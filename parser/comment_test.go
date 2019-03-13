package parser_test

import (
	"testing"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func TestParser_Comment(t *testing.T) {
	t.Parallel()

	tt := []ParserTest{
		{
			Name:       "It should parse comments",
			Expression: "// A test comment",
			ExpectedNode: &ast.Comment{
				Token: token.New(" A test comment", token.COMMENT, 0, 0),
				Value: "A test comment",
			},
		},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}
