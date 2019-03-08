package parser

import (
	"github.com/davidsbond/dave/ast"
)

func (p *Parser) parseReturnStatement() ast.Node {
	stmt := &ast.ReturnStatement{
		Token: p.currentToken,
	}

	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)

	return stmt
}
