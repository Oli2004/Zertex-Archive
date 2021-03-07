package tamper

import (
	"strings"
)

func Space2Comment(query string) string {
	return strings.ReplaceAll(query, " ", "/**/")
}