package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateCallExpression(node *ast.CallExpression, scope *object.Scope) object.Object {
	fn := Evaluate(node.Function, scope)

	if isError(fn) {
		return nil
	}

	var args []object.Object

	for _, arg := range node.Arguments {
		val := Evaluate(arg, scope)

		if isError(val) {
			return nil
		}

		args = append(args, val)
	}

	switch function := fn.(type) {
	default:
		return object.Error("expected function, got %s", function.Type())
	case nil:
		return object.Error("expected function, got nil")
	case *object.Function:
		return evaluateSynchronousFunction(function, args, scope)
	}
}

func evaluateSynchronousFunction(fn *object.Function, args []object.Object, scope *object.Scope) object.Object {
	newScope := scope.NewChildScope()

	for i, param := range fn.Parameters {
		newScope.Set(param.Value, args[i])
	}

	result := Evaluate(fn.Body, newScope)

	switch obj := result.(type) {
	case *object.ReturnValue:
		return obj.Value
	default:
		return obj
	}
}
