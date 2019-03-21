package evaluator

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/object"
)

func ifExpression(node *ast.IfExpression, scope *object.Scope) object.Object {
	var result object.Object
	condition := Evaluate(node.Condition, scope)

	if isError(condition) {
		return condition
	}

	if condition == TRUE {
		result = Evaluate(node.Consequence, scope.NewChildScope())
	}

	if condition == FALSE && node.Alternative != nil {
		result = Evaluate(node.Alternative, scope.NewChildScope())
	}

	return result
}
