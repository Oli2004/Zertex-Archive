package postgres

var Queries = map[string]string {
	"test": "",
	"current_db": "",
	"db_name": "",
	"db_count": "",
	"current_version": "",
	"current_user": "",
	"table_count": "",
	"table_name": "",
	"column_count": "", // table, database
	"column_name": "",
	"row_count": "",
	"dump_column": "", // column, db, table, ord, index
	"dump_multi_column": "", // column, db, table, ord, index
	"column_from_table": "",
	"table_from_column": "",
	"table_count_with_column": "",
}

var Flags = []string {
	"pg_query()",
	"PgSQL",
	"pg_fetch_assoc()",
	"pg_free_result()",
	"pg_num_rows()",
}

var SystemDBNames = []string{"information_schema", "pg_catalog", "pg_toast", "pgagent"}