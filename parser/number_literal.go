package parser

import (
	"fmt"
	"strconv"

	"github.com/davidsbond/dave/ast"
)

func (p *Parser) parseNumberLiteral() ast.Node {
	lit := &ast.NumberLiteral{Token: p.currentToken}

	value, err := strconv.ParseFloat(p.currentToken.Literal, 64)

	if err != nil {
		msg := fmt.Errorf("could not parse %q as number", p.currentToken.Literal)
		p.Errors = append(p.Errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}
