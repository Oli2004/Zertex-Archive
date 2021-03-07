package tamper

func AppendNullByte(query string) string {
	return query + "%00"
}