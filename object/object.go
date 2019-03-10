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

	// The Hashable interfaces defines a type that can be stored
	// in a hash
	Hashable interface {
		HashKey() HashKey
	}

	// The HashKey type is used to store unique keys in a hash map
	HashKey struct {
		Type  Type
		Value float64
	}
)
