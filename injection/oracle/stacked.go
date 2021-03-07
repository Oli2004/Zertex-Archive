package oracle

var S_Vectors = map[string]string {
	"dbms_pipe": "%s;SELECT DBMS_PIPE.RECEIVE_MESSAGE((%s),5) FROM DUAL%s",

}