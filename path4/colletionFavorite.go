package main

import "fmt"

type Favorite struct {
	Title []string
}

func Add(favorite *Favorite, title string) {
	favorite.Title = append(favorite.Title, title)
}

func PrintAll(favorite *Favorite) []string {
	return favorite.Title
}

func PrintUnique(favorite *Favorite) []string {
	seen := make(map[string]struct{}) // создаём пустую map
	unique := []string{}

	for _, title := range favorite.Title {
		if _, ok := seen[title]; ok {
			continue // такое название уже было, пропускаем
		}
		seen[title] = struct{}{} // запоминаем название
		unique = append(unique, title)
	}
	return unique
}

func main() {
	fav := &Favorite{}
	Add(fav, "Dark")
	Add(fav, "Breaking Bad")
	Add(fav, "Dark")

	summary := struct {
		Owner      string
		ItemsCount int
	}{
		Owner:      "Alice",
		ItemsCount: len(PrintUnique(fav)),
	}

	fmt.Println(summary)         // {Alice 2}
	fmt.Printf("%+v\n", summary) // {Owner:Alice ItemsCount:2}
	fmt.Println(summary.Owner, summary.ItemsCount)
}
