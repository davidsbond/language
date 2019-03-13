package evaluator

import (
	"math"

	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func prefixExpression(node *ast.PrefixExpression, scope *object.Scope) object.Object {
	right := Evaluate(node.Right, scope)

	if isError(right) {
		return right
	}

	switch node.Operator {
	case token.BANG:
		return bangOperatorExpression(right)
	case token.MINUS:
		return minusPrefixOperatorExpression(right)
	case token.SQRT:
		return sqrtPrefixOperatorExpression(right)
	default:
		return object.Error("type %s does not support operator %s", right.Type(), node.Operator)
	}
}

func sqrtPrefixOperatorExpression(right object.Object) object.Object {
	switch val := right.(type) {
	case *object.Number:
		return &object.Number{Value: math.Sqrt(val.Value)}
	case *object.Atomic:
		return sqrtPrefixOperatorExpression(val.Value())
	case *object.Constant:
		return sqrtPrefixOperatorExpression(val.Value)
	default:
		return object.Error("type %s does not support operator '√'", right.Type())
	}
}

func bangOperatorExpression(right object.Object) object.Object {
	switch val := right.(type) {
	case *object.Boolean:
		if val == TRUE {
			return FALSE
		}

		return TRUE
	case *object.Atomic:
		return bangOperatorExpression(val.Value())
	case *object.Constant:
		return bangOperatorExpression(val.Value)
	default:
		return object.Error("type %s does not support operator '!'", right.Type())
	}
}

func minusPrefixOperatorExpression(right object.Object) object.Object {
	switch val := right.(type) {
	case *object.Number:
		return &object.Number{Value: -val.Value}
	case *object.Constant:
		return minusPrefixOperatorExpression(val.Value)
	case *object.Atomic:
		return minusPrefixOperatorExpression(val.Value())
	default:
		return object.Error("type %s does not support operator '-'", right.Type())
	}
}
