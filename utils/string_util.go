package utils

import (
	"regexp"
	"strings"
)

func FormatPrice(text string) string {
	text = strings.Replace(text, "8%", "", -1)
	text = strings.Replace(text, "10%", "", -1)
	text = regexp.MustCompile(`[^a-z0-9, ]+`).ReplaceAllString(text, "")
	text = strings.Replace(text, " ", "", -1)
	return text
}
