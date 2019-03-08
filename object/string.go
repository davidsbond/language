package object

const (
	// TypeString is the type for a string value.
	TypeString = "String"
)

type (
	// The String type represents a string value in memory.
	String struct {
		Value string
	}
)

// Type returns the type of the object.
func (str *String) Type() Type {
	return TypeString
}

// Clone creates a copy of the current object that can be used
// without modifying the original value
func (str *String) Clone() Object {
	return &String{Value: str.Value}
}

func (str *String) String() string {
	return str.Value
}
