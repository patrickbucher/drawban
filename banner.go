package drawban

import (
	"fmt"
	"strings"
)

// DrawBanner draws a border around the content with the given border rune, and
// returns the banner's lines. A horizontal inner padding of one character is
// added, but no vertical padding.
func DrawBanner(content []string, border rune) []string {
	if len(content) == 0 {
		return []string{}
	}

	trimmedLines := trim(content)
	longestLine := maxlen(trimmedLines)
	if longestLine == 0 {
		return []string{}
	}

	bannerWidth := longestLine + 4 // border and padding, left and right
	bar := strings.Repeat(string(border), bannerWidth)
	wrappedLines := make([]string, len(trimmedLines)+2 /* add bars */)
	wrappedLines[0] = bar
	for i, line := range trimmedLines {
		lineLen := len(line)
		line += strings.Repeat(" ", longestLine-lineLen)
		wrappedLines[i+1] = fmt.Sprintf("%c %s %c", border, line, border)
	}
	wrappedLines[len(wrappedLines)-1] = bar
	return wrappedLines
}

func maxlen(lines []string) int {
	max := -1
	for _, line := range lines {
		if l := len(line); l > max {
			max = l
		}
	}
	return max
}

func trim(lines []string) []string {
	trimmed := make([]string, len(lines))
	for i := range lines {
		trimmed[i] = strings.TrimSpace(lines[i])
	}
	return trimmed
}
