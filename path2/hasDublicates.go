package main

import "fmt"

func hasDuplicates(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 1}
	fmt.Println(hasDuplicates(nums)) // Output: true
}
