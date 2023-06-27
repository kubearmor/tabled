/*
Copyright 2023 Rahul Jadhav.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package json2table convert input json to table format
package drawtable

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nyrahul/tabled/pkg/config"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func getRow(rec []string) table.Row {
	var row table.Row
	for _, x := range rec {
		row = append(row, x)
	}
	return row
}

func Csv2Table(cfg config.Config) {
	records := readCsvFile(cfg.InFile)
	if len(records) <= 1 {
		log.Fatal("insufficient entries in file " + cfg.InFile)
	}

	t := table.NewWriter()
	if cfg.Caption != "" {
		t.SetCaption(cfg.Caption)
	}
	if cfg.Title != "" {
		t.SetTitle(cfg.Title)
	}
	t.SetOutputMirror(os.Stdout)
	for idx, rec := range records {
		if idx == 0 {
			t.AppendHeader(getRow(rec))
		} else {
			t.AppendRow(getRow(rec))
		}
	}
	t.SetStyle(table.StyleLight)
	//	t.SetStyle(table.StyleColoredBright)
	t.Render()
	// t.RenderMarkdown()
}
