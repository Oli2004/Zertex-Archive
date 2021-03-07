package tamper

import b64 "encoding/base64"

func B64Encode(query string) string {
	return b64.URLEncoding.EncodeToString([]byte(query))
}