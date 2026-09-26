package main

import "fmt"

var catalog = map[string]float64{
	"Star Wars": 8.6,
	"Avatar":    7.8,
	"Inception": 8.8,
	"matrix":    8.7,
}

func addFilm(catalog map[string]float64, title string, raiting int) {
	catalog[title] = float64(raiting)
}

func topFilm(catalog map[string]float64, min float64) []string {
	var topFilms []string
	for title, rating := range catalog {
		if rating >= min {
			topFilms = append(topFilms, title)
		}
	}
	return topFilms
}

func avgRating(catalog map[string]float64) float64 {
	var sum float64
	for _, rating := range catalog {
		sum += rating
	}
	return sum / float64(len(catalog))
}

func main() {
	addFilm(catalog, "The Godfather", 9)
	fmt.Println(catalog)
	addFilm(catalog, "The Godfather", 9)
	fmt.Println(catalog)
	topFilms := topFilm(catalog, 8.5)
	fmt.Println(topFilms)
	average := avgRating(catalog)
	fmt.Printf("Average rating: %.2f\n", average)
}
