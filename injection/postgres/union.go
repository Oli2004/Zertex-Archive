package postgres

var U_Vectors = map[string]string {
	"single": "%s UNION SELECT %s %s",
	"multi": "%s UNION ALL SELECT %s %s",
}