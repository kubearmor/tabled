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
	"errors"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/exp/slices"

	//"github.com/jedib0t/go-pretty/v6/text"
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

func colorNameToEnum(cstr string) (text.Color, error) {
	offset := text.Reset
	stroffset := 0
	var colors []string
	colors = []string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}
	if strings.HasPrefix(cstr, "FgHi") {
		offset = text.FgHiBlack
		stroffset = 4
	} else if strings.HasPrefix(cstr, "Fg") {
		offset = text.FgBlack
		stroffset = 2
	} else if strings.HasPrefix(cstr, "BgHi") {
		offset = text.BgHiBlack
		stroffset = 4
	} else if strings.HasPrefix(cstr, "Bg") {
		offset = text.BgBlack
		stroffset = 2
	} else {
		colors = []string{"Reset", "Bold", "Faint", "Italic", "Underline", "BlinkSlow",
			"BlinkRapid", "ReverseVideo", "Concealed", "CrossedOut"}
	}
	cidx := slices.Index(colors, cstr[stroffset:])
	if cidx < 0 {
		return offset, errors.New("invalid color")
	}
	return offset + text.Color(cidx), nil
}

func getTextColors(paint config.Paint) text.Colors {
	var colors text.Colors
	for _, color := range paint.Color {
		c, err := colorNameToEnum(color)
		if err != nil {
			log.Printf("invalid color <%s>", color)
			continue
		}
		colors = append(colors, c)
	}
	return colors
}

func getAlign(align string) text.Align {
	if align == "" {
		return text.AlignDefault
	}
	alignSet := []string{"default", "left", "center", "justify", "right"}
	i := slices.Index(alignSet, align)
	if i < 0 {
		return text.AlignDefault
	}
	return text.Align(i)
}

func getColConfig(col config.ColConfig) (table.ColumnConfig, error) {
	var colcfg table.ColumnConfig
	colcfg.Name = col.Name
	colcfg.Align = getAlign(col.Align)
	colcfg.AlignHeader = colcfg.Align
	if col.MaxWidth > 0 {
		colcfg.WidthMax = col.MaxWidth
	}
	colcfg.AutoMerge = col.AutoMerge
	if col.MinWidth > 0 {
		colcfg.WidthMin = col.MinWidth
	}
	return colcfg, nil
}

func getAllColConfigs(hdr []string, cols []config.ColConfig) []table.ColumnConfig {
	var colcfgs []table.ColumnConfig
	for _, col := range cols {
		idx := slices.Index(hdr, col.Name)
		if idx < 0 {
			continue
		}
		c, err := getColConfig(col)
		if err == nil {
			colcfgs = append(colcfgs, c)
		}
	}
	return colcfgs
}

func Csv2Table(cfg config.Config) {
	records := readCsvFile(cfg.InFile)
	if len(records) <= 1 {
		log.Fatal("insufficient entries in file " + cfg.InFile)
	}

	t := table.NewWriter()
	if cfg.YamlCfg.Table.Caption != "" {
		t.SetCaption(cfg.YamlCfg.Table.Caption)
	}
	if cfg.YamlCfg.Table.Title != "" {
		t.SetTitle(cfg.YamlCfg.Table.Title)
	}
	t.SetOutputMirror(os.Stdout)
	var hdr []string
	for idx, rec := range records {
		if idx == 0 {
			hdr = rec
			t.AppendHeader(getRow(rec))
		} else {
			t.AppendRow(getRow(rec))
		}
	}
	colcfgs := getAllColConfigs(hdr, cfg.YamlCfg.Columns)
	if len(colcfgs) > 0 {
		t.SetColumnConfigs(colcfgs)
	}
	t.SetRowPainter(func(row table.Row) text.Colors {
		for _, col := range cfg.YamlCfg.Columns {
			idx := slices.Index(hdr, col.Name)
			if idx < 0 {
				continue
			}
			for _, p := range col.Paint {
				colors := getTextColors(p)
				if len(colors) <= 0 || p.Regex == "" {
					continue
				}
				if colval, ok := row[idx].(string); ok {
					match, _ := regexp.MatchString(p.Regex, colval)
					if match {
						return colors
					}
				}
			}
		}
		return nil
	})
	t.SetStyle(table.StyleLight)
	t.Render()
	// t.RenderMarkdown()
}
