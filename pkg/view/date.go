package view

import (
	"strings"
)

func Date(digits string, focused bool) string {
	template := "dd.mm.yyyy"
	var res strings.Builder

	visualCursor := len(digits)
	if len(digits) >= 2 {
		visualCursor++
	}
	if len(digits) >= 4 {
		visualCursor++
	}

	digitIdx := 0
	for i := 0; i < len(template); i++ {
		isDot := template[i] == '.'
		isCursor := focused && i == visualCursor

		var char string
		if isDot {
			char = "."
		} else if digitIdx < len(digits) {
			char = string(digits[digitIdx])
			digitIdx++
		} else {
			char = string(template[i])
		}

		if isCursor {
			res.WriteString("\033[7m" + char + "\033[0m")
		} else if isDot {
			res.WriteString(char)
		} else if digitIdx > 0 && i < visualCursor {
			res.WriteString(char)
		} else {
			res.WriteString("\033[4;2m" + char + "\033[0m")
		}
	}

	if focused && visualCursor == len(template) {
		res.WriteString("\033[7m \033[0m")
	}

	return res.String()
}
