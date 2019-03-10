package evaluator

import "github.com/davidsbond/dave/object"

func evaluatePrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evaluateBangOperatorExpression(right)
	// case "-":
	// 	return evaluateMinusPrefixOperatorExpression(right)
	default:
		return object.Error("type %s does not support operator %s", right.Type(), operator)
	}
}

func evaluateBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}
