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

func loadCatalog(files []string) (Movies []Movie) {
	for _, file := range files {
		Movies = append(Movies, readJSONcatlog(file)...)
	}
	return Movies
}

func readJSONcatlog(file string) []Movie {
	defer trackTime(time.Now(), "readJSONcatlog")
	time.Sleep(80 * time.Microsecond)
	jsonFile, err := os.Open(file)
	fmt.Println("Successfully Opened %s", file)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return nil
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	var films []Movie

	err = json.Unmarshal(byteValue, &films)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return films
}

func main() {
	files := []string{"./path3/JSON/movies_valid.json", "./path3/JSON/movies_broken.json", "./path3/JSON/movies_empty.json"}
	fmt.Println(loadCatalog(files))
}
