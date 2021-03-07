package tamper

import (
	"github.com/dlclark/regexp2"
	"strings"
	"fmt"
)

var reg *regexp2.Regexp

func init() {
	reg = regexp2.MustCompile(`(?i)(\b(AND|OR)\b\s+)([^>]+?)\s*>\s*(\w+|'[^']+')`, 0)
}

func Greatest(query string) string {
	retVal := query
	if match, _ := reg.FindStringMatch(query); match != nil {
		gps := match.Groups()
		retVal = strings.ReplaceAll(query, match.String(), fmt.Sprintf("%sGREATEST(%s,%s+1)=%s", gps[1].String(), gps[3].String(), gps[4].String(), gps[3].String()))
	}
	return retVal
}