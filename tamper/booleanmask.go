package tamper

import "strings"

func BooleanMask(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(query, "AND", "&&"), "OR", "||")
}