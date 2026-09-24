package main

import "fmt"

func filterByRating(rating []float64, min float64) []float64 {
	result := make([]float64, 0, len(rating))
	for _, r := range rating {
		if r >= min {
			result = append(result, r)
		}
	}
	return result
}

func main() {
	ratings := []float64{4.5, 3.0, 4.0, 5.0, 3.5}
	minRating := 4.0
	result := filterByRating(ratings, minRating)
	fmt.Println(result)
}
