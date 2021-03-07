package mssql

import (
	"fmt"
	"strings"
)

var Queries = map[string]string {
	"test": "",
	"current_user": "",
	"current_version": "",
	"current_db": "",
	"db_count": "",
	"db_name": "",
	"table_count": "",
	"table_name": "",
	"column_count": "",
	"column_name": "",
	"row_count": "",
	"dump_column": "",
	"dump_multi_column": "",
	"column_from_table": "",
	"table_from_column": "",
	"table_count_with_column": "",
}

var Flags = []string {
	"SQL Server",
	"Microsoft OLE DB Provider",
	"ODBC Driver",
	"SQLExecDirect",
}

func ConvertStringToChars(payload string) string {
	charArr := make([]string, len(payload))
	for index, char := range payload {
		charArr[index] = fmt.Sprintf("CHAR(%d)", char)
	}
	return strings.Join(charArr, "+")
}
var SystemDBNames = []string{"Northwind", "master", "model", "msdb", "pubs", "tempdb", "Resource", "ReportServer", "ReportServerTempDB"}