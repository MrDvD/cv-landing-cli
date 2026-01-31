package view

import "cv-landing-cli/pkg/tool/color"

func Textfield(content string, cursor int, isFocused bool) string {
	val := content
	isPlaceholder := false
	if val == "" {
		val = "type there..."
		isPlaceholder = true
	}
	runes := []rune(val)

	line := color.StyledLine{}

	for i := range runes {
		attrs := []color.Attribute{}

		if isPlaceholder {
			attrs = append(attrs, color.AttrGray)
		}

		if isFocused && i == cursor {
			attrs = append(attrs, color.AttrInvert)
		}

		line.Segments = append(line.Segments, color.Segment{
			Text:       string(runes[i]),
			Attributes: attrs,
		})
	}

	if isFocused && cursor >= len(runes) {
		attrs := []color.Attribute{}
		if isPlaceholder {
			attrs = append(attrs, color.AttrGray)
		}
		attrs = append(attrs, color.AttrInvert)

		line.Segments = append(line.Segments, color.Segment{
			Text:       " ",
			Attributes: attrs,
		})
	}

	return line.Render()
}
