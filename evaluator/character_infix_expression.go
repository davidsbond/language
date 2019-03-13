package evaluator

import (
	"fmt"

	"github.com/davidsbond/dave/object"
	"github.com/davidsbond/dave/token"
)

func characterInfixExpression(operator string, left, right object.Object) object.Object {
	// Get the *object.Character instances for both the left and right objects. Handle
	// atomic/constants appropriately.
	trueLeft, err := getCharacterFromObject(left)

	if err != nil {
		return object.Error(err.Error())
	}

	trueRight, err := getCharacterFromObject(right)

	if err != nil {
		return object.Error(err.Error())
	}

	switch operator {
	default:
		return object.Error("type %s does not support operator %s", trueLeft.Type(), operator)
	// 'a' + 'b'
	case token.PLUS:
		return &object.String{Value: string([]rune{trueLeft.Value, trueRight.Value})}
	// 'a' - 'b'
	case token.MINUS:
		return &object.Character{Value: trueLeft.Value - trueRight.Value}
	// 'a' * 'b'
	case token.ASTERISK:
		return &object.Character{Value: trueLeft.Value * trueRight.Value}
	// 'a' / 'b'
	case token.SLASH:
		return &object.Character{Value: trueLeft.Value / trueRight.Value}
	// 'a' > 'b'
	case token.LT:
		if trueLeft.Value < trueRight.Value {
			return TRUE
		}

		return FALSE
	// 'a' > 'b'
	case token.GT:
		if trueLeft.Value > trueRight.Value {
			return TRUE
		}

		return FALSE
	// 'a' == 'b'
	case token.EQUALS:
		if trueLeft.Value == trueRight.Value {
			return TRUE
		}

		return FALSE
	// 'a' != 'b'
	case token.NOTEQ:
		if trueLeft.Value != trueRight.Value {
			return TRUE
		}

		return FALSE
	}
}

func getCharacterFromObject(obj object.Object) (*object.Character, error) {
	switch val := obj.(type) {
	case *object.Constant:
		return getCharacterFromObject(val.Value)
	case *object.Atomic:
		return getCharacterFromObject(val.Value())
	case *object.Character:
		return val, nil
	case nil:
		return nil, fmt.Errorf("cannot cast nil value")
	default:
		return nil, fmt.Errorf("cannot cast type %s to type Character", obj.Type())
	}
}
