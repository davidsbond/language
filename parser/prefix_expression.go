package parser

import (
	"github.com/davidsbond/language/ast"
)

func (p *Parser) parsePrefixExpression() ast.Node {
	expression := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}
