package object

import "hash/fnv"

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

// HashKey creates a unique key for this value for use in hash maps.
func (str *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(str.Value))

	return HashKey{Type: str.Type(), Value: float64(h.Sum64())}
}

func (str *String) String() string {
	return str.Value
}
