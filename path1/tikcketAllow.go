package main

import "fmt"

func isAllowed(age int, allowAge int) bool {
	return age >= allowAge
}

func main() {
	age := 25
	allowAge := 18
	allowed := isAllowed(age, allowAge)
	fmt.Println(allowed)
}
