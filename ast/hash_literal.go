package ast

import (
	"strings"

	"github.com/davidsbond/language/token"
)

type (
	// The HashLiteral type represents a literal hash object in source code.
	// For example:
	// const hash = {
	//   "a": "b",
	//	 "b": "c"
	// }
	HashLiteral struct {
		Token *token.Token
		Pairs map[Node]Node
	}
)

func (hl *HashLiteral) String() string {
	var out strings.Builder

	out.WriteString("{\n")

	i := 0
	for key, val := range hl.Pairs {
		out.WriteRune('\t')
		out.WriteString(key.String())
		out.WriteString(": ")
		out.WriteString(val.String())

		if i != len(hl.Pairs)-1 {
			out.WriteString(",")
		}

		out.WriteRune('\n')
		i++
	}

	out.WriteString("}\n")

	return out.String()
}
