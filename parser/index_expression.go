package parser

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/token"
)

func (p *Parser) parseIndexExpression(left ast.Node) ast.Node {
	exp := &ast.IndexExpression{
		Token: p.currentToken,
		Left:  left,
	}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}
