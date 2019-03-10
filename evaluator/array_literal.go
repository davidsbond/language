package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateArrayLiteral(node *ast.ArrayLiteral, scope *object.Scope) object.Object {
	var elements []object.Object

	for _, elem := range node.Elements {
		val := Evaluate(elem, scope)

		if isError(val) {
			return val
		}

		elements = append(elements, val)
	}

	return &object.Array{Elements: elements}
}
