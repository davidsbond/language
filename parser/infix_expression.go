package parser

import "github.com/davidsbond/language/ast"

func (p *Parser) parseInfixExpression(left ast.Node) ast.Node {
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()

	p.nextToken()

	expression.Right = p.parseExpression(precedence)

	return expression
}
