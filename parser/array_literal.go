package parser

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func (p *Parser) parseArrayLiteral() ast.Node {
	array := &ast.ArrayLiteral{Token: p.currentToken}

	array.Elements = p.parseExpressionList(token.RBRACKET)

	return array
}
