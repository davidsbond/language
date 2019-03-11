package parser

import "github.com/davidsbond/dave/ast"

func (p *Parser) parsePostfixExpression(left ast.Node) ast.Node {
	expression := &ast.PostfixExpression{
		Token: p.currentToken,
	}

	if ident, ok := left.(*ast.Identifier); ok {
		expression.Left = ident
	} else {
		p.error("expected an identifier, got %s", left.String())
		return nil
	}

	p.nextToken()

	expression.Operator = p.currentToken.Literal

	return expression
}
