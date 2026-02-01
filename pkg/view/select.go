package view

import (
	"fmt"
	"strings"
)

func Select(options []string, cursor int, selected int, focused bool) string {
	var b strings.Builder

	for i, option := range options {
		checked := "[ ]"
		if i == selected {
			checked = "[x]"
		}

		line := fmt.Sprintf("%s %s", checked, option)

		if i == cursor && focused {
			fmt.Fprintf(&b, "\033[1;36m%s\033[0m", line)
		} else {
			b.WriteString(line)
		}
		if i < len(options)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}
