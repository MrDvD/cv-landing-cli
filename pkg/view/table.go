package view

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Column struct {
	Header string
	Weight int
}

type TableConfig[T any] struct {
	Columns   []Column
	RowMapper func(item T) []string
}

func Table[T any](data []T, config TableConfig[T]) {
	if len(data) == 0 {
		return
	}

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}

	totalWeight := 0
	for _, col := range config.Columns {
		totalWeight += col.Weight
	}

	usableWidth := width - (len(config.Columns) + 1)
	colWidths := make([]int, len(config.Columns))
	sumWidths := 0

	for i, col := range config.Columns {
		colWidths[i] = (col.Weight * usableWidth) / totalWeight
		sumWidths += colWidths[i]
	}

	gap := usableWidth - sumWidths
	for i := range gap {
		colWidths[i%len(colWidths)]++
	}

	line := strings.Repeat("-", width)

	fmt.Println(line)
	for i, col := range config.Columns {
		fmt.Printf("| %-*s ", colWidths[i]-2, truncate(col.Header, colWidths[i]-2))
	}
	fmt.Println("|")
	fmt.Println(line)

	for _, item := range data {
		cells := config.RowMapper(item)
		for i, content := range cells {
			if i < len(colWidths) {
				fmt.Printf("| %-*s ", colWidths[i]-2, truncate(content, colWidths[i]-2))
			}
		}
		fmt.Println("|")
	}
	fmt.Println(line)
}

func truncate(s string, limit int) string {
	runes := []rune(s)
	if len(runes) <= limit {
		return s
	}
	if limit < 3 {
		return string(runes[:limit])
	}
	return string(runes[:limit-3]) + "..."
}

func StrPtr(s *string) string {
	if s == nil {
		return "-"
	}
	return *s
}
