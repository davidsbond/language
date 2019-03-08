package object

type (
	// The Constant type is a wrapper around Object that indicates the
	// object is immutable.
	Constant struct {
		Value Object
	}
)

// Type returns the type of the underlying object.
func (c *Constant) Type() Type {
	return c.Value.Type()
}

// Clone creates a copy of the current object that can be used
// without modifying the original value
func (c *Constant) Clone() Object {
	return c.Value.Clone()
}

func (c *Constant) String() string {
	return c.Value.String()
}
