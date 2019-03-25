package evaluator

import (
	"fmt"
	"math"

	"github.com/davidsbond/language/object"
	"github.com/davidsbond/language/token"
)

func numberInfixExpression(operator string, left, right object.Object) object.Object {
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
	// 1 + 1
	case token.PLUS:
		return &object.Number{Value: trueLeft.Value + trueRight.Value}
	// 1 - 1
	case token.MINUS:
		return &object.Number{Value: trueLeft.Value - trueRight.Value}
	// 1 * 1
	case token.ASTERISK:
		return &object.Number{Value: trueLeft.Value * trueRight.Value}
	// 1 / 1
	case token.SLASH:
		return &object.Number{Value: trueLeft.Value / trueRight.Value}
	// 1 % 1
	case token.MOD:
		return &object.Number{Value: math.Mod(trueLeft.Value, trueRight.Value)}
	// 1 ^ 1
	case token.POW:
		return &object.Number{Value: math.Pow(trueLeft.Value, trueRight.Value)}
	// 1 < 1
	case token.LT:
		if trueLeft.Value < trueRight.Value {
			return TRUE
		}

		return FALSE
	// 1 > 1
	case token.GT:
		if trueLeft.Value > trueRight.Value {
			return TRUE
		}

		return FALSE
	// 1 == 1
	case token.EQUALS:
		if trueLeft.Value == trueRight.Value {
			return TRUE
		}

		return FALSE
	// 1 != 1
	case token.NOTEQ:
		if trueLeft.Value != trueRight.Value {
			return TRUE
		}

		return FALSE
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
