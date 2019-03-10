package object

import "hash/fnv"

const (
	// TypeCharacter is the type for a character value.
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

// HashKey creates a unique key for this value for use in hash maps.
func (c *Character) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(string(c.Value)))

	return HashKey{Type: c.Type(), Value: float64(h.Sum64())}
}

func (c *Character) String() string {
	return string(c.Value)
}
