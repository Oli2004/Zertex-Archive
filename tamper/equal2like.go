package tamper

import "strings"

func Equal2Like(query string) string {
	return strings.ReplaceAll(query, "=", " LIKE ")
}