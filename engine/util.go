package engine

import (
	"strconv"
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

// FilterInvalidChar ...
func FilterInvalidChar(content string) string {
	for _, v := range []string{`"`, `|`, "`"} {
		content = strings.ReplaceAll(content, v, " ")
	}
	return content
}

// ConvertUnicode utf-8 to unicode
func ConvertUnicode(content string) string {
	textQuoted := strconv.QuoteToASCII(content)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return textUnquoted
}
