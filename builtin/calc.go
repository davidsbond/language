package builtin

import (
	"math"

	"github.com/davidsbond/dave/object"
)

// Sin provides the ability to calculate the sine of the first radian
// argument.
func Sin(args ...object.Object) object.Object {
	return numberOperation("sin", math.Sin, args...)
}

// Tan provides the ability to calculate the tangent of the first radian
// argument.
func Tan(args ...object.Object) object.Object {
	return numberOperation("tan", math.Tan, args...)
}

// Cos provides the ability to calculate the cosine of the first radian
// argument.
func Cos(args ...object.Object) object.Object {
	return numberOperation("cos", math.Cos, args...)
}

// Log provides the ability to calculate the natural logarithm of the
// first argument
func Log(args ...object.Object) object.Object {
	return numberOperation("log", math.Log, args...)
}

// Sqrt provides the ability to calculate the square root of a given number.
func Sqrt(args ...object.Object) object.Object {
	return numberOperation("sqrt", math.Sqrt, args...)
}

// Ceil rounds up a given number to the least integer value greather than it.
func Ceil(args ...object.Object) object.Object {
	return numberOperation("ceil", math.Ceil, args...)
}

// Floor rounds down a given number to the greatest integer value less than it.
func Floor(args ...object.Object) object.Object {
	return numberOperation("floor", math.Floor, args...)
}

func numberOperation(name string, fn func(float64) float64, args ...object.Object) object.Object {
	if len(args) > 1 || len(args) == 0 {
		return object.Error("built-in '%s' function only takes one argument(s)", name)
	}

	switch obj := args[0].(type) {
	default:
		return object.Error("built-in '%s' does not support type %s", name, obj.Type())
	case *object.Number:
		return &object.Number{Value: fn(obj.Value)}
	case *object.Constant:
		return numberOperation(name, fn, obj.Value)
	case *object.Atomic:
		return numberOperation(name, fn, obj.Value())
	}
}
