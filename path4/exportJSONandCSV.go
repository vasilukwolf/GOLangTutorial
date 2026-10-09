package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

type Exporter interface {
	Export() string
}

type CSVExport struct {
	FileName string
	Data     [][]string // строки таблицы
}

func (c CSVExport) Export() string {
	var sb strings.Builder
	w := csv.NewWriter(&sb)
	w.WriteAll(c.Data) // записывает все строки и сбрасывает буфер
	return sb.String()
}

type JSONExport struct {
	FileName string
	Data     any // любая структура, срез или map
}

func (j JSONExport) Export() string {
	b, err := json.MarshalIndent(j.Data, "", "  ")
	if err != nil {
		return "ошибка: " + err.Error()
	}
	return string(b)
}

func RunExport(e Exporter) string {
	return e.Export()
}

func main() {
	// 1. Литерал с именами полей (самый частый и понятный способ)
	csvExp := CSVExport{
		FileName: "episodes.csv",
		Data: [][]string{
			{"title", "season", "duration"},
			{"Pilot", "1", "45"},
			{"Second", "1", "50"},
		},
	}

	jsonExp := JSONExport{
		FileName: "episode.json",
	}

	// 2. Указатель на структуру через &
	csvPtr := &CSVExport{FileName: "report.csv", Data: [][]string{{"a", "b"}}}

	// 3. Пустой объект с нулевыми значениями, поля заполняются позже
	var jsonEmpty JSONExport
	jsonEmpty.FileName = "empty.json"
	jsonEmpty.Data = []string{"one", "two"}

	fmt.Println(RunExport(csvExp))
	fmt.Println(RunExport(jsonExp))
	fmt.Println(RunExport(csvPtr))
	fmt.Println(RunExport(jsonEmpty))
}
