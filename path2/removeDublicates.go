package main

import "fmt"

func removeDuplicates(s []string) []string {
	result := []string{}
	contains := func(slice []string, str string) bool {
		for _, v := range slice {
			if v == str {
				return true
			}
		}
		return false
	}
	for _, str := range s {
		if !contains(result, str) {
			result = append(result, str)
		}
	}
	return result
}

func main() {
	strs := []string{"apple", "banana", "apple", "orange", "banana"}
	result := removeDuplicates(strs)
	fmt.Println(result)
	strings := []string{"a", "a", "and", "a", "b", "b"}
	result2 := removeDuplicates(strings)
	fmt.Println(result2)
	emptystrings := []string{}
	result3 := removeDuplicates(emptystrings)
	fmt.Println(result3)
	nilstrings := []string(nil)
	result4 := removeDuplicates(nilstrings)
	fmt.Println(result4)
}
