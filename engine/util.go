package engine

import (
	"strings"
	"unicode/utf8"
)

// FilterEmoji ...
func FilterEmoji(content string) string {
	var res []string
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			res = append(res, string(value))
		}
	}
	return strings.Join(res, "")
}
