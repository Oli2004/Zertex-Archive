package oracle

import (
	"fmt"
	"strings"
)

var Queries = map[string]string {
	"test": "",
	"current_user": "",
	"db_count": "",
	"db_name": "",
	"current_db": "",
	"current_version": "",
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
	"ORA-01756:",
	"ociparse()",
	"ocifreestatement()",
	"ocifetchstatement()",
}

func ConvertStringToChars(payload string) string {
	charArr := make([]string, len(payload))
	for index, char := range payload {
		charArr[index] = fmt.Sprintf("CHR(%d)", char)
	}
	return strings.Join(charArr, "||")
}

var SystemDBNames = []string{"ADAMS", "ANONYMOUS", "APEX_030200", "APEX_PUBLIC_USER", "APPQOSSYS", "AURORA$ORB$UNAUTHENTICATED", "AWR_STAGE", "BI", "BLAKE", "CLARK", "CSMIG", "CTXSYS", "DBSNMP", "DEMO", "DIP", "DMSYS", "DSSYS", "EXFSYS", "FLOWS_%", "FLOWS_FILES", "HR", "IX", "JONES", "LBACSYS", "MDDATA", "MDSYS", "MGMT_VIEW", "OC", "OE", "OLAPSYS", "ORACLE_OCM", "ORDDATA", "ORDPLUGINS", "ORDSYS", "OUTLN", "OWBSYS", "PAPER", "PERFSTAT", "PM", "SCOTT", "SH", "SI_INFORMTN_SCHEMA", "SPATIAL_CSW_ADMIN_USR", "SPATIAL_WFS_ADMIN_USR", "SYS", "SYSMAN", "SYSTEM", "TRACESVR", "TSMSYS", "WK_TEST", "WKPROXY", "WKSYS", "WMSYS", "XDB", "XS$NULL"}