package evaluator

import (
	"math"

	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func evaluateNumberInfixExpression(operator string, left, right object.Object) object.Object {
	// Get the *object.Number instances for both the left and right objects. Handle
	// atomic/constants appropriately.
	trueLeft := getNumberFromObject(left)
	trueRight := getNumberFromObject(right)

	switch operator {
	default:
		return nil
	case token.PLUS:
		return &object.Number{Value: trueLeft.Value + trueRight.Value}
	case token.MINUS:
		return &object.Number{Value: trueLeft.Value - trueRight.Value}
	case token.ASTERISK:
		return &object.Number{Value: trueLeft.Value * trueRight.Value}
	case token.SLASH:
		return &object.Number{Value: trueLeft.Value / trueRight.Value}
	case token.MOD:
		return &object.Number{Value: math.Mod(trueLeft.Value, trueRight.Value)}
	case token.LT:
		return &object.Boolean{Value: trueLeft.Value < trueRight.Value}
	case token.GT:
		return &object.Boolean{Value: trueLeft.Value > trueRight.Value}
	case token.EQUALS:
		return &object.Boolean{Value: trueLeft.Value == trueRight.Value}
	}
}

func getNumberFromObject(obj object.Object) *object.Number {
	switch val := obj.(type) {
	case *object.Constant:
		return val.Value.(*object.Number)
	case *object.Atomic:
		return val.Value().(*object.Number)
	default:
		return obj.(*object.Number)
	}
}
