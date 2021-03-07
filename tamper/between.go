package tamper

import (
	regexp2 "github.com/dlclark/regexp2"
	"strings"
	"fmt"
)

var first *regexp2.Regexp
var second *regexp2.Regexp
var third *regexp2.Regexp

func init() {
	first = regexp2.MustCompile(`(?i)(\b(AND|OR)\b\s+)(?!.*\b(AND|OR)\b)([^>]+?)\s*>\s*([^>]+)\s*\Z`, 0)
	second = regexp2.MustCompile(`\s*>\s*(\d+|'[^']+'|\w+\(\d+\))`, 0)
	third = regexp2.MustCompile(`(?i)(\b(AND|OR)\b\s+)(?!.*\b(AND|OR)\b)([^=]+?)\s*=\s*([\w()]+)\s*`, 0)
}

func Between(query string) string {
	retVal := query
	if match, _ := first.FindStringMatch(query); match != nil {
		gps := match.Groups()
		
		retVal = strings.ReplaceAll(retVal, match.String(), fmt.Sprintf("%s %s NOT BETWEEN 0 AND %s", gps[2].Captures[0].String(), gps[4].Captures[0].String(), gps[5].Captures[0].String()))
	} else if match, _ := second.FindStringMatch(query); match != nil {
		gps := match.Groups()
		retVal = strings.ReplaceAll(query, match.String(), fmt.Sprintf(" NOT BETWEEN 0 AND %s", gps[1].String()))
	}
	
	if retVal == query {
		if match, _ := third.FindStringMatch(query); match != nil {
			gps := match.Groups()
			retVal = strings.ReplaceAll(retVal, match.String(), fmt.Sprintf("%s %s BETWEEN %s AND %s", gps[2].Captures[0].String(), gps[4].Captures[0].String(), gps[5].Captures[0].String(), gps[5].Captures[0].String()))
		}
	}
	return strings.ReplaceAll(query, "'", "%00%27")
}