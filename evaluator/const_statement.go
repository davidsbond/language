package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateConstStatement(node *ast.ConstStatement, scope *object.Scope) object.Object {
	val := Evaluate(node.Value, scope)

	if isError(val) {
		return val
	}

	return scope.Set(node.Name.Value, &object.Constant{Value: val})
}
