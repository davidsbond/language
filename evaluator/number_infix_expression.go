package evaluator

import (
	"fmt"
	"math"

	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func evaluateNumberInfixExpression(operator string, left, right object.Object) object.Object {
	// Get the *object.Number instances for both the left and right objects. Handle
	// atomic/constants appropriately.
	trueLeft, err := getNumberFromObject(left)

	if err != nil {
		return object.Error(err.Error())
	}

	trueRight, err := getNumberFromObject(right)

	if err != nil {
		return object.Error(err.Error())
	}

	switch operator {
	default:
		return object.Error("type Number does not support operator %s", operator)
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

func getNumberFromObject(obj object.Object) (*object.Number, error) {
	switch val := obj.(type) {
	case *object.Constant:
		return val.Value.(*object.Number), nil
	case *object.Atomic:
		return val.Value().(*object.Number), nil
	case *object.Number:
		return obj.(*object.Number), nil
	case nil:
		return nil, fmt.Errorf("cannot cast nil value")
	default:
		return nil, fmt.Errorf("cannot cast type %s to type Number", obj.Type())
	}
}
