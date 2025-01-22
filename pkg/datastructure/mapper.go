package datastructure

import (
	"iter"
	"maps"
)

// GetFruitNames returns the names of all fruits in the FruitNumberMap.
func GetFruitNames() iter.Seq[string] {
	return maps.Keys(FruitNumberMap)
}
