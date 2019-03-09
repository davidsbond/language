package object

const (
	// TypeNull is the type for null objects.
	TypeNull = "null"
)

type (
	// The Null type represents an unset object.
	Null struct {
	}
)

// Type returns the object's type.
func (n *Null) Type() Type {
	return TypeNull
}

// Clone returns null
func (n *Null) Clone() Object {
	return n
}

func (n *Null) String() string {
	return TypeNull
}
