package tamper

import (
	"time"
	"math/rand"
	"strings"
	"fmt"
)

// cloudflare + cloudbric

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func LuaNGINX(query string) string {
	if !strings.Contains(query, "?") {
		return query
	}
	parts := strings.Split(query, "?")
	retVal := parts[0] + "?"
	for i:=0; i < 110; i++ {
		retVal += fmt.Sprintf("%s=%s&", StringWithCharset(seededRand.Intn(3), charset), StringWithCharset(seededRand.Intn(3), charset))
	}
	return retVal + parts[1]
}