package mysql

var S_Vectors = map[string]string {
	"basic": "%s;SELECT %s%s",
	"query": "%s;(SELECT * FROM (SELECT(%s))CWTq)%s",
}