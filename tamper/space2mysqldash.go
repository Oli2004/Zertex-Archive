package tamper

import "strings"

func Space2MySQLDash(query string) string {
	return strings.ReplaceAll(query, " ", "–%0A")
}