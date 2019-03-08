package object

const (
	// TypeCharacter is the node type for a character object.
	TypeCharacter = "Character"
)

type (
	// The Character type represents a character value in memory.
	Character struct {
		Value rune
	}
)

// Type returns the type of the object.
func (c *Character) Type() Type {
	return TypeCharacter
}

// Clone creates a copy of the current object that can be used
// without modifying the original value
func (c *Character) Clone() Object {
	return &Character{Value: c.Value}
}

func (c *Character) String() string {
	return string(c.Value)
}
