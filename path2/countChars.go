package main

import "fmt"

func countChars(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	return counts
}

func main() {
	s := "hello world"
	result := countChars(s)
	for char, count := range result {
		fmt.Printf("%c: %d\n", char, count)
	}
}
