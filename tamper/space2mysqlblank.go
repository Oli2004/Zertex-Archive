package tamper

import (
	"strings"
	"math/rand"
)

var blanks = []string{"%2B", "%0D", "%0C"}

func Space2MySQLBlank(query string) string {
	return strings.ReplaceAll(query, " ", blanks[rand.Intn(2)])
}