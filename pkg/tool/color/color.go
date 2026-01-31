package color

import "strings"

type Attribute int

const (
	AttrGray Attribute = iota
	AttrInvert
	AttrBold
)

var startCodes = map[Attribute]string{
	AttrGray:   "\x1b[90m",
	AttrInvert: "\x1b[7m",
	AttrBold:   "\x1b[1m",
}

type Segment struct {
	Text       string
	Attributes []Attribute
}

type StyledLine struct {
	Segments []Segment
}

func (l StyledLine) Render() string {
	var sb strings.Builder
	for _, seg := range l.Segments {
		for _, attr := range seg.Attributes {
			sb.WriteString(startCodes[attr])
		}
		sb.WriteString(seg.Text)
		sb.WriteString("\x1b[0m")
	}
	return sb.String()
}
