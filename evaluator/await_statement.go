package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateAwaitStatement(node *ast.AwaitStatement, scope *object.Scope) object.Object {
	switch fn := node.Value.(type) {
	default:
		return object.Error("await keyword can only be used on function calls")
	case *ast.CallExpression:
		fn.Awaited = true

		return Evaluate(fn, scope)
	}
}
