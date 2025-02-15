package datastructure

import (
	"iter"
	"maps"
)

var fruitNumberMap map[string]int = map[string]int{
	"apple":  1,
	"banana": 2,
	"cherry": 3,
}

// GetFruitNames returns the names of all fruits in the fruitNumberMap.
func GetFruitNames() iter.Seq[string] {
	return maps.Keys(fruitNumberMap)
}

// FruitNameExists checks if a name exists in the fruitNumberMap.
func FruitNameExists(fruit string) bool {
	_, ok := fruitNumberMap[fruit]
	return ok
}
