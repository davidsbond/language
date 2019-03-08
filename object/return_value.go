package object

const (
	// TypeReturnValue is the object type for `return` values.
	TypeReturnValue = "ReturnValue"
)

type (
	// The ReturnValue type represents a return value in the source code.
	// For example return 1
	ReturnValue struct {
		Value Object
	}
)

// Type returns the type of the object.
func (rv *ReturnValue) Type() Type {
	return TypeReturnValue
}

// Clone creates a copy of the current object that can be used
// without modifying the original value
func (rv *ReturnValue) Clone() Object {
	return rv.Value.Clone()
}

func (rv *ReturnValue) String() string {
	return rv.Value.String()
}
