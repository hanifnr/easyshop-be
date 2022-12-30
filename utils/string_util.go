package utils

import (
	"regexp"
	"strings"
)

func FormatPrice(text string) string {
	text = strings.Replace(text, "10%", "", -1)
	return regexp.MustCompile(`[^a-zA-Z0-9 ,]+`).ReplaceAllString(text, "")
}
