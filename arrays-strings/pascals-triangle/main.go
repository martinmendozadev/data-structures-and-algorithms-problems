// O(n)

package main

import "fmt"

func main() {
	const row uint64 = 62
	array := []uint64{1}
	var (
		aux			uint64 = 0
		coefficient uint64 = 1
		xExponent	uint64 = row
		yExponent	uint64 = 1
	)

	for xExponent > 0 {
		aux = (coefficient * xExponent) / yExponent
		array = append(array, aux)
		coefficient = aux
		xExponent--
		yExponent++
	}

	fmt.Println(array)
}
