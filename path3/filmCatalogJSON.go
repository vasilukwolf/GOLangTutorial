package main

import (
	"fmt"
	"os"
	"time"
)

type Movie struct {
	Name string `json:"Name"`
	Year int `json:"Year"`
}

func trackTime(start time.Time,name string){
	fmt.Printf("%s: %v\n", name,time.Since(start).Round(time.Millisecond) )
}

func loadCatalog(files []string) (movies []Movie, err error) {

	for _, file range files{
		os.OpenFile("file", os.O_CREATE, os.ModePerm)
	    defer file.Close()
	}

	encoder := decoder
	encoder.Encode(makeData(1000000))

	return nil, err
}
