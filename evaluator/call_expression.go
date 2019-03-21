package evaluator

import (
	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/object"
)

func callExpression(node *ast.CallExpression, scope *object.Scope) object.Object {
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

	switch fnc := fn.(type) {
	default:
		return object.Error("expected function, got %s", fnc.Type())
	case nil:
		return object.Error("expected function, got nil")
	case *object.Function:
		if fnc.Async && node.Awaited {
			res := make(chan object.Object, 1)

			go func() {
				res <- function(fnc, args, scope)
				close(res)
			}()

			return <-res
		}

		if fnc.Async {
			go function(fnc, args, scope)
			return NULL
		}

		return function(fnc, args, scope)

	case object.Builtin:
		return fnc(args...)
	}
}

func function(fn *object.Function, args []object.Object, scope *object.Scope) object.Object {
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
