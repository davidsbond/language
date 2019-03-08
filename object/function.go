package object

import (
	"strings"

	"github.com/davidsbond/dave/ast"
)

const (
	TypeFunction = "Function"
)

type (
	Function struct {
		Name       *ast.Identifier
		Parameters []*ast.Identifier
		Body       *ast.BlockStatement
		Scope      *Scope
	}
)

func (fn *Function) Type() Type {
	return TypeFunction
}

func (fn *Function) Clone() Object {
	return &Function{
		Name:       fn.Name,
		Parameters: fn.Parameters,
		Body:       fn.Body,
		Scope:      fn.Scope,
	}
}

func (fn *Function) String() string {
	var out strings.Builder

	out.WriteString("function ")
	out.WriteString(fn.Name.String())
	out.WriteString("(")

	for i, ident := range fn.Parameters {
		out.WriteString(ident.String())

		if i != len(fn.Parameters)-1 {
			out.WriteString(",")
		}
	}

	out.WriteString(") {\n")
	out.WriteString(fn.Body.String())
	out.WriteString("\n}")

	return out.String()
}
