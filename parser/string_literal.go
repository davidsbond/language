package parser

import (
	"github.com/davidsbond/language/ast"
)

func (p *Parser) parseStringLiteral() ast.Node {
	return &ast.StringLiteral{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}
