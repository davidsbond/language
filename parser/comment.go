package parser

import (
	"strings"

	"github.com/davidsbond/language/ast"
)

func (p *Parser) parseComment() ast.Node {
	return &ast.Comment{
		Token: p.currentToken,
		Value: strings.Trim(p.currentToken.Literal, " "),
	}
}
