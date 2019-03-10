package object

import (
	"strings"
)

const (
	TypeHash = "Hash"
)

type (
	Hash struct {
		Pairs map[HashKey]HashPair
	}

	HashPair struct {
		Key   Object
		Value Object
	}
)

func (h *Hash) Type() Type {
	return TypeHash
}

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
