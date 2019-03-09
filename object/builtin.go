package object

type (
	// The Builtin type represents a built-in function that can be called
	// from the source code.
	Builtin func(args ...Object) Object
)

// Type returns the object's type.
func (bi Builtin) Type() Type {
	return TypeFunction
}

// Clone creates a copy of the current object that can be used
// without modifying the original value
func (bi Builtin) Clone() Object {
	return bi
}

func (bi Builtin) String() string {
	return "built-in function"
}
