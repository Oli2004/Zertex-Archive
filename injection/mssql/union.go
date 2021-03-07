package mssql

var U_Vectors = map[string]string {
	"single": "%s UNION SELECT %s %s",
	"all": "%s UNION ALL SELECT %s %s",
}