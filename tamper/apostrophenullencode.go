package tamper

import "strings"

func ApostropheNullEncode(query string) string {
	return strings.ReplaceAll(query, "'", "%00%27")
}