package drawban

import (
	"strings"
	"testing"
)

type drawBannerTest struct {
	Content  []string
	Border   rune
	Expected []string
}

var tests = []drawBannerTest{
	{
		// no content -> no banner
		[]string{}, '*', []string{},
	},
	{
		// only whitespace content -> no banner
		[]string{" "}, '*', []string{},
	},
	{
		// single line -> banner around the line
		[]string{"foo"},
		'*',
		[]string{
			"*******",
			"* foo *",
			"*******",
		},
	},
	{
		// some other border for variation
		[]string{"foo"},
		'#',
		[]string{
			"#######",
			"# foo #",
			"#######",
		},
	},
	{
		// multiple lines -> multi-line banner
		[]string{"This", "is", "a", "test"},
		'*',
		[]string{
			"********",
			"* This *",
			"* is   *",
			"* a    *",
			"* test *",
			"********",
		},
	},
	{
		// long line -> sets banner width
		[]string{"This", "is", "a rather nasty", "test"},
		'*',
		[]string{
			"******************",
			"* This           *",
			"* is             *",
			"* a rather nasty *",
			"* test           *",
			"******************",
		},
	},
}

func TestDrawBanner(t *testing.T) {
	for _, test := range tests {
		actual := DrawBanner(test.Content, test.Border)
		if !equal(actual, test.Expected) {
			expectedStr := strings.Join(test.Expected, "\n")
			actualStr := strings.Join(actual, "\n")
			t.Errorf("DrawBanner for content %v and border '%c'\n\nexpected:\n%s\n\ngot:\n%s\n",
				test.Content, test.Border, expectedStr, actualStr)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
