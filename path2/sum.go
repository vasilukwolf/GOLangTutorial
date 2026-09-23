package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println(result)
	nums = []int{}
	fmt.Println(sum(nums...))
	nums = nil
	fmt.Println(sum(nums...))
}
