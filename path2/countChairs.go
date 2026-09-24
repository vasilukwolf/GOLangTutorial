package main

import "fmt"

func countChairs(s []string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	return counts
}

func main() {
	s := "hello world"
	result := countChairs(s)
	fmt.Println(result)
}
