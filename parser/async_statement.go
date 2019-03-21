package parser

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/token"
)

func (p *Parser) parseAsyncStatement() ast.Node {
	stmt := &ast.AsyncStatement{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.FUNCTION) {
		return nil
	}

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}
