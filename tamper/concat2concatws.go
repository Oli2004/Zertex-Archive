package tamper

import "strings"

func Concat2ConcatWS(query string) string {
	return strings.ReplaceAll(query, "CONCAT(", "CONCAT_WS(MID(CHAR(0),0,0),")
}