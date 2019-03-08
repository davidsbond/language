package parser

import (
	"github.com/davidsbond/dave/ast"
)

func (p *Parser) parseStringLiteral() ast.Node {
	return &ast.StringLiteral{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}
