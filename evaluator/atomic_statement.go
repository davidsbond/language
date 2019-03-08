package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateAtomicStatement(node *ast.AtomicStatement, scope *object.Scope) object.Object {
	val := Evaluate(node.Value, scope)

	// if isError(val) {
	// 	return val
	// }

	scope.Set(node.Name.Value, object.MakeAtomic(val))
	return nil
}
