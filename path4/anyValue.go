package main

import "fmt"

type SearchResult struct {
	Title string
	URL   string
}

func InspectSearchResult(value any) {
	str, ok := value.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("value is not a string")
	}

	v, ok := value.(int)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("value is not an int")
	}

	switch v := value.(type) {
	case int:
		fmt.Printf("%d is an int\n", v)
	case string:
		fmt.Printf("%s is a string\n", v)
	case bool:
		fmt.Printf("%t is a bool\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	InspectSearchResult("Hello, World!")
	InspectSearchResult(42)
	InspectSearchResult(true)
	InspectSearchResult(SearchResult{Title: "Go Programming", URL: "https://golang.org"})
	InspectSearchResult([]string{"Анна", "Борис", "Виктор"})
}
