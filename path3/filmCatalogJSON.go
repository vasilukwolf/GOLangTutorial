package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Movie struct {
	Name string `json:"Name"`
	Year int    `json:"Year"`
}

func trackTime(start time.Time, name string) {
	fmt.Printf("%s: %v\n", name, time.Since(start).Round(time.Millisecond))
}

func loadCatalog(files []string) ([]Movie, error) {
	defer trackTime(time.Now(), "loadCatalog")
	var all []Movie
	for _, file := range files {
		movies, err := loadFile(file)
		if err != nil {
			return nil, err
		}
		all = append(all, movies...)
	}
	return all, nil
}

func loadFile(file string) ([]Movie, error) {
	defer trackTime(time.Now(), file)
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", file, err)
	}
	defer jsonFile.Close()
	fmt.Printf("Successfully opened %s\n", file)

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", file, err)
	}

	var films []Movie
	err = json.Unmarshal(byteValue, &films)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", file, err)
	}

	return films, nil
}

func main() {
	files := []string{"./path3/JSON/movies_valid.json", "./path3/JSON/movies_broken.json", "./path3/JSON/movies_empty.json"}
	loadCatalog(files)
}
