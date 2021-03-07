package postgres

import (
	"fmt"
	"strings"
)

var E_Vectors = map[string]string {
	"where": "",
	"param_repl": "",
	"param_repl_gen": "",
	"orderby": "",
	"orderby_gen": "",
	"groupby": "",
	"groupby_gen": "",
}

func ConvertStringToChars(payload string) string {
	charArr := make([]string, len(payload))
	for index, char := range payload {
		charArr[index] = fmt.Sprintf("CHR(%d)", char)
	}
	return strings.Join(charArr, "||")
}