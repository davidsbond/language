package parser

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func (p *Parser) parseBooleanLiteral() ast.Node {
	return &ast.BooleanLiteral{
		Token: p.currentToken,
		Value: p.currentToken.Type == token.TRUE,
	}
}
