package evaluator

import (
	"fmt"

	"github.com/davidsbond/language/object"
	"github.com/davidsbond/language/token"
)

func booleanInfixExpression(operator string, left, right object.Object) object.Object {
	// Get the *object.Boolean instances for both the left and right objects. Handle
	// atomic/constants appropriately.
	trueLeft, err := getBooleanFromObject(left)

	if err != nil {
		return object.Error(err.Error())
	}

	trueRight, err := getBooleanFromObject(right)

	if err != nil {
		return object.Error(err.Error())
	}

	switch operator {
	default:
		return object.Error("type %s does not support operator %s", trueLeft.Type(), operator)
	case token.EQUALS:
		if trueLeft.Value == trueRight.Value {
			return TRUE
		}

		return FALSE
	case token.NOTEQ:
		if trueLeft.Value != trueRight.Value {
			return TRUE
		}

		return FALSE
	}
}

func getBooleanFromObject(obj object.Object) (*object.Boolean, error) {
	switch val := obj.(type) {
	case *object.Constant:
		return getBooleanFromObject(val.Value)
	case *object.Atomic:
		return getBooleanFromObject(val.Value())
	case *object.Boolean:
		return val, nil
	case nil:
		return nil, fmt.Errorf("cannot cast nil value")
	default:
		return nil, fmt.Errorf("cannot cast type %s to type Boolean", obj.Type())
	}
}
