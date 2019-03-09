package ast

import "github.com/davidsbond/dave/token"

type (
	HashLiteral struct {
		Token *token.Token
		Pairs map[Node]Node
	}
)
