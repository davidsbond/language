package object

import (
	"strings"
)

const (
	// TypeHash is the type returned by hash objects.
	TypeHash = "Hash"
)

type (
	// The Hash type represents a hash stored in memory. Each key is uniquely generated
	// based on the type used as the underlying key.
	Hash struct {
		Pairs map[HashKey]HashPair
	}

	// The HashPair type represents a key/value pair stored in a hash.
	HashPair struct {
		Key   Object
		Value Object
	}
)

// Type returns this object's type.
func (h *Hash) Type() Type {
	return TypeHash
}

// Clone creates a copy of the map.
func (h *Hash) Clone() Object {
	hsh := &Hash{
		Pairs: make(map[HashKey]HashPair),
	}

	for key, pair := range h.Pairs {
		hsh.Pairs[key] = HashPair{
			pair.Key.Clone(),
			pair.Value.Clone(),
		}
	}

	return hsh
}

func (h *Hash) String() string {
	var out strings.Builder

	out.WriteString("{ ")

	i := 0
	for _, pair := range h.Pairs {
		out.WriteString(pair.Key.String())
		out.WriteString(": ")
		out.WriteString(pair.Value.String())

		if i != len(h.Pairs)-1 {
			out.WriteString(", ")
		}

		i++
	}

	out.WriteString(" }")
	return out.String()
}
