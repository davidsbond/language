package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func evaluatePrefixExpression(node *ast.PrefixExpression, scope *object.Scope) object.Object {
	right := Evaluate(node.Right, scope)

	if isError(right) {
		return right
	}

	switch node.Operator {
	case token.BANG:
		return evaluateBangOperatorExpression(right)
	case token.MINUS:
		return evaluateMinusPrefixOperatorExpression(right)
	default:
		return object.Error("type %s does not support operator %s", right.Type(), node.Operator)
	}
}

func evaluateBangOperatorExpression(right object.Object) object.Object {
	switch val := right.(type) {
	case *object.Boolean:
		if val == TRUE {
			return FALSE
		}

		return TRUE
	case *object.Atomic:
		return evaluateBangOperatorExpression(val.Value())
	case *object.Constant:
		return evaluateBangOperatorExpression(val.Value)
	default:
		return object.Error("type %s does not support operator '!'", right.Type())
	}
}

func evaluateMinusPrefixOperatorExpression(right object.Object) object.Object {
	switch val := right.(type) {
	case *object.Number:
		return &object.Number{Value: -val.Value}
	case *object.Constant:
		return evaluateMinusPrefixOperatorExpression(val.Value)
	case *object.Atomic:
		return evaluateMinusPrefixOperatorExpression(val.Value())
	default:
		return object.Error("type %s does not support operator '-'", right.Type())
	}
}
