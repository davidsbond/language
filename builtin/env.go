package builtin

import (
	"os"

	"github.com/davidsbond/dave/object"
)

// SetEnv sets an environment variable. The first argument is the key and second is the value. The first
// argument must be a string and the second argument will be converted to a string.
func SetEnv(args ...object.Object) object.Object {
	if len(args) > 2 || len(args) == 0 {
		return object.Error("built-in 'set_env' function only takes two arguments")
	}

	// Iterate over the first argument's type to ensure
	// it is a string.
	var key string
	switch obj := args[0].(type) {
	default:
		return object.Error("built-in 'set_env' does not support type %s", obj.Type())
	case *object.String:
		key = obj.Value

	// If we have a constant or atomic value passed in, obtain its value and
	// use recursion to check for their stringiness.
	case *object.Constant:
		return SetEnv(obj.Value, args[1])
	case *object.Atomic:
		return SetEnv(obj.Value(), args[1])
	}

	// Once we have a string key, we convert the value to a string and set the
	// env var.
	if err := os.Setenv(key, args[1].String()); err != nil {
		return object.Error(err.Error())
	}

	return &object.Null{}
}

// GetEnv gets an environment variable. The first and only argument is the key which must be
// a string.
func GetEnv(args ...object.Object) object.Object {
	if len(args) > 1 || len(args) == 0 {
		return object.Error("built-in 'get_env' function only takes two arguments")
	}

	// Iterate over the first argument's type to ensure
	// it is a string.
	var key string
	switch obj := args[0].(type) {
	default:
		return object.Error("built-in 'get_env' does not support type %s", obj.Type())
	case *object.String:
		key = obj.Value

	// If we have a constant or atomic value passed in, obtain its value and
	// use recursion to check for their stringiness.
	case *object.Constant:
		return GetEnv(obj.Value, args[1])
	case *object.Atomic:
		return GetEnv(obj.Value(), args[1])
	}

	// Return a new string with the env var value.
	return &object.String{Value: os.Getenv(key)}
}
