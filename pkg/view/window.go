package view

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

func visualWidth(s string) int {
	stripped := ansiRegex.ReplaceAllString(s, "")
	return utf8.RuneCountInString(stripped)
}

func Window(width, height int, category string, content string, tips string) string {
	if width < 5 || height < 5 {
		return "Terminal too small"
	}

	var builder strings.Builder

	headerPrefix := "── " + category + " "
	vWidthHeader := visualWidth(headerPrefix)

	builder.WriteString(headerPrefix)
	if width > vWidthHeader {
		builder.WriteString(strings.Repeat("─", width-vWidthHeader))
	}
	builder.WriteString("\n")

	contentLines := strings.Split(content, "\n")
	bodyHeight := height - 2

	for i := 0; i < bodyHeight; i++ {
		line := ""
		if i < len(contentLines) {
			line = contentLines[i]
		}

		vWidthLine := visualWidth(line)

		builder.WriteString(line)

		paddingCount := width - vWidthLine
		if paddingCount > 0 {
			builder.WriteString(strings.Repeat(" ", paddingCount))
		}
		builder.WriteString("\n")
	}

	footerText := " " + tips + " "
	vWidthFooter := visualWidth(footerText)

	paddingFooter := width - vWidthFooter
	if paddingFooter > 0 {
		builder.WriteString(strings.Repeat("─", paddingFooter))
	}
	builder.WriteString(footerText)

	return builder.String()
}
