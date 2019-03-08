// Package object contains types that represent in-memory objects while the interpreter
// is evaluating source code.
package object

type (
	// The Object interface defines behavior for all objects.
	Object interface {
		Type() Type
		String() string
		Clone() Object
	}

	// The Type type contains an object's type.
	Type string
)
