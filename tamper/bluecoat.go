package tamper

import "strings"

func Bluecoat(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(query, " ", "%09"), "=", " LIKE ")
}