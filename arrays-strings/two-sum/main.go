// Time complexity: O(n)

package main

import "fmt"

func main() {
	hashMap := map[int]int{}
	nums := []int {3, -1, 0, 1, 4}
	target := 0

	for i, v := range nums {
		complement := target - v
		index, ok := hashMap[complement]

		if ok {
			fmt.Printf("[%v, %v]\n", index, i)
			break
		}
		
		hashMap[v] = i
	}
	fmt.Println(hashMap)
}
