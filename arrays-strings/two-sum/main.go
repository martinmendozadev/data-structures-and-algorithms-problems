// Time complexity: O(n^2)

package main

import "fmt"

func main() {
	nums := []int {2, 7, 11, 15}
	target := 9

	for i, v := range nums {
		for j := 1; j < len(nums); j++ {
			if v + nums[j] == target {
				fmt.Printf("[%v, %v]\n", i, j)
				break
			}
		}
	}
}
