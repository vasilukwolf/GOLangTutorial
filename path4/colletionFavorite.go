package main

import "fmt"

type Favorite struct {
	Title []string
}

func (f *Favorite) Add(title string) {
	f.Title = append(f.Title, title)
}

func (f Favorite) PrintAll() {
	fmt.Println("All favorite titles:")
	for _, title := range f.Title {
		fmt.Println(title)
	}
}

func (f Favorite) PrintUnique() {
	seen := make(map[string]struct{}) // создаём пустую map
	unique := []string{}

	for _, title := range f.Title {
		if _, ok := seen[title]; ok {
			continue // такое название уже было, пропускаем
		}
		seen[title] = struct{}{} // запоминаем название
		unique = append(unique, title)
	}

	fmt.Println("Unique favorite titles:")
	for _, title := range unique {
		fmt.Println(title)
	}
}

func main() {
	fav := &Favorite{}
	fav.Add("Dark")
	fav.Add("Breaking Bad")
	fav.Add("Dark")

	summary := struct {
		Owner      string
		ItemsCount int
	}{
		Owner:      "Alice",
		ItemsCount: len(fav.Title),
	}

	fmt.Println(summary)         // {Alice 2}
	fmt.Printf("%+v\n", summary) // {Owner:Alice ItemsCount:2}
	fmt.Println(summary.Owner, summary.ItemsCount)
}
