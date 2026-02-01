package view

import (
	"cv-landing-cli/pkg/tool/color"
	"strings"
)

func Textfield(content string, cursor int, isFocused bool, width, height int) string {
	innerWidth := max(width-2, 0)
	val := content
	isPlaceholder := false
	if val == "" {
		val = "type here..."
		isPlaceholder = true
	}
	runes := []rune(val)

	cursorRow := 0
	if innerWidth > 0 {
		cursorRow = cursor / innerWidth
	}
	scrollOffset := 0
	if cursorRow >= height {
		scrollOffset = cursorRow - height + 1
	}

	var output strings.Builder
	sideBorder := "│"

	for r := range height {
		line := color.StyledLine{}
		line.Segments = append(line.Segments, color.Segment{Text: sideBorder})

		logicalRow := r + scrollOffset

		for c := range innerWidth {
			realIndex := (logicalRow * innerWidth) + c

			attrs := []color.Attribute{}
			char := " "

			if realIndex < len(runes) {
				char = string(runes[realIndex])
				if isPlaceholder {
					attrs = append(attrs, color.AttrGray)
				}
			}
			if isFocused && realIndex == cursor {
				attrs = append(attrs, color.AttrInvert)
			}

			line.Segments = append(line.Segments, color.Segment{
				Text:       char,
				Attributes: attrs,
			})
		}

		line.Segments = append(line.Segments, color.Segment{Text: sideBorder})

		if r > 0 {
			output.WriteString("\n")
		}
		output.WriteString(line.Render())
	}

	return output.String()
}
