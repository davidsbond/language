package evaluator

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/object"
)

func awaitStatement(node *ast.AwaitStatement, scope *object.Scope) object.Object {
	switch fn := node.Value.(type) {
	default:
		return object.Error("await keyword can only be used on function calls")
	case *ast.CallExpression:
		fn.Awaited = true

		return Evaluate(fn, scope)
	}
}
