package main

import(
	"fmt"
)

func main() {
	testScores := map[string]float32 {
		"Harry": 87.3,
		"Hermione": 105,
		"Ronald": 63.5,
		"Neville": 27,
	}

	c := clone(testScores)
	
	fmt.Println(c)
}

// since the keys have to be comparable in maps therefor you can't use "any" because slices are not comparable. So instead us "comparable" to make the keys comparable.
func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}