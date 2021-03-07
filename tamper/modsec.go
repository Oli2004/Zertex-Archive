package tamper

import "fmt"

func ModSec(query string) string {
	return fmt.Sprintf("/*!00000%s*/", query)
}