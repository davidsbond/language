package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
)

func evaluateCallExpression(node *ast.CallExpression, scope *object.Scope) object.Object {
	fn := Evaluate(node.Function, scope)

	if isError(fn) {
		return fn
	}

	var args []object.Object

	for _, arg := range node.Arguments {
		val := Evaluate(arg, scope)

		if isError(val) {
			return val
		}

		args = append(args, val)
	}

	switch function := fn.(type) {
	default:
		return object.Error("expected function, got %s", function.Type())
	case nil:
		return object.Error("expected function, got nil")
	case *object.Function:
		if function.Async && node.Awaited {
			res := make(chan object.Object, 1)

			go func() {
				res <- evaluateFunction(function, args, scope)
				close(res)
			}()

			return <-res
		}

		if function.Async {
			go evaluateFunction(function, args, scope)
			return &object.Null{}
		}

		return evaluateFunction(function, args, scope)

	case object.Builtin:
		return function(args...)
	}
}

func evaluateFunction(fn *object.Function, args []object.Object, scope *object.Scope) object.Object {
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
