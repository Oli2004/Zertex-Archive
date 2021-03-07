package tamper

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandomCase(query string) string {
	var retVal string
	for _, c := range query {
		switch rand.Intn(1) {
		case 0:
			retVal += strings.ToUpper(fmt.Sprintf("%c", c))
		case 1:
			retVal += strings.ToLower(fmt.Sprintf("%c", c))
		}
	}
	return retVal
}