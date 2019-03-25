package parser

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func (p *Parser) parseBooleanLiteral() ast.Node {
	return &ast.BooleanLiteral{
		Token: p.currentToken,
		Value: p.currentToken.Type == token.TRUE,
	}
}
