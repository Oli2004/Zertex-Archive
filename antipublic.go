package modules

import (
	"bufio"
	"database/sql"
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/qtui"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/paulbellamy/ratecounter"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"strings"
	"time"
)

type AntiPublicModule struct {
	Index                  int
	Public                 int
	Private                int
	DB                     *sqlx.DB
	CheckPublicUrlQuery    *sql.Stmt
	CheckPublicDomainQuery *sql.Stmt
	InsertUrlQuery         *sql.Stmt
	InsertDomainQuery      *sql.Stmt
}

const (
	URLS = iota
	DOMAINS
)

var AntiPublic *AntiPublicModule

func NewAntiPublic() *AntiPublicModule {
	database, _ := sqlx.Connect("sqlite3", "./ap.db")
	database.Exec("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY, url TEXT UNIQUE)")
	database.Exec("CREATE TABLE IF NOT EXISTS domains (id INTEGER PRIMARY KEY, domain TEXT UNIQUE)")

	checkurlStmt, err := database.Prepare("SELECT COUNT(*) FROM urls WHERE url = ?")
	if err != nil {
		panic(err.Error())
	}
	checkdomainStmt, err := database.Prepare("SELECT COUNT(*) FROM domains WHERE domain = ?")
	if err != nil {
		panic(err.Error())
	}
	inserturlStmt, err := database.Prepare("INSERT OR IGNORE INTO urls (url) VALUES (?)")
	if err != nil {
		panic(err.Error())
	}
	insertdomainStmt, err := database.Prepare("INSERT OR IGNORE INTO domains (domain) VALUES (?)")
	if err != nil {
		panic(err.Error())
	}

	return &AntiPublicModule{
		Index:               0,
		DB:                  database,
		CheckPublicUrlQuery: checkurlStmt,
		InsertUrlQuery:      inserturlStmt,
		CheckPublicDomainQuery: checkdomainStmt,
		InsertDomainQuery: insertdomainStmt,
	}
}

func (ap *AntiPublicModule) GetSelected() int {
	if qtui.Main.AntipubDomainMode.IsChecked() {
		return DOMAINS
	}
	return URLS
}

func (ap *AntiPublicModule) GetTable() string {
	switch ap.GetSelected() {
	case DOMAINS:
		return "domains"
	case URLS:
		return "urls"
	}
	return ""
}

func (ap *AntiPublicModule) GetColumn() string {
	switch ap.GetSelected() {
	case DOMAINS:
		return "domain"
	case URLS:
		return "url"
	}
	return ""
}

func (ap *AntiPublicModule) LoadToDB(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return ap.InsertLines(lines)
}

func (ap *AntiPublicModule) Count(table string) int {
	var c int
	row := ap.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table))
	if err := row.Scan(&c); err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return c
}

func (ap *AntiPublicModule) Size() int64 {
	stat, err := os.Stat("ap.db")
	if err != nil {
		return -1
	}

	return stat.Size()
}

func (ap *AntiPublicModule) InsertLines(lines []string) (int, error) {
	var a []string
	for _, u := range lines {
		_u, err := url.Parse(u)
		if err != nil {
			continue
		}
		if ap.GetSelected() == DOMAINS {
			u = _u.Hostname()
		} else {
			u = _u.String()
		}
		
		a = append(a, fmt.Sprintf("(\"%s\")", []byte(u)))
	}
	var query string
	switch ap.GetSelected() {
	case DOMAINS:
		query = fmt.Sprintf("INSERT OR IGNORE INTO domains (domain) VALUES %s", strings.Join(a, ","))
	case URLS:
		query = fmt.Sprintf("INSERT OR IGNORE INTO urls (url) VALUES %s", strings.Join(a, ","))
	}
	//fmt.Println(query)
	res, err := ap.DB.Exec(query)
	if err != nil {
		return 0, err
	}

	aff, _ := res.RowsAffected()
	return int(aff), nil
}

func (ap *AntiPublicModule) InsertLine(line string) error {
	var err error
	switch ap.GetSelected() {
	case DOMAINS:
		_u, err := url.Parse(line)
		if err != nil {
			return err
		}
		_, err = ap.InsertDomainQuery.Exec(_u.Hostname())
	case URLS:
		_, err = ap.InsertUrlQuery.Exec(line)
	}

	return err
}

func (ap *AntiPublicModule) CheckPublic(u string) bool {
	var a int
	var err error
	switch ap.GetSelected() {
	case DOMAINS:
		_u, err := url.Parse(u)
		if err != nil {
			return false
		}
		row := ap.CheckPublicDomainQuery.QueryRow(_u.Hostname())
		err = row.Scan(&a)
	case URLS:
		row := ap.CheckPublicUrlQuery.QueryRow(u)
		err = row.Scan(&a)
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	return a == 1
}

func (ap *AntiPublicModule) Start(urls []string) {
	utils.GlobalIndex = 0
	utils.ErrorCounter = 0
	utils.RateCounter = ratecounter.NewRateCounter(time.Second)

	ap.Public=0
	ap.Private=0

	done := make(chan interface{})
	defer close(done)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-utils.Done:
				return
			case <-time.After(250 * time.Millisecond):
				qtui.Main.AntipubPublic.Display2(ap.Public)
				qtui.Main.AntipubPrivate.Display2(ap.Private)
				qtui.Main.AntipubPrivateRatio.Display(fmt.Sprintf("%.2f%%", float64(ap.Private)/float64(ap.Index)))
			}
		}
	}()

	outputCh := make(chan string)

	go func() {
		f := utils.CreateFileTimeStamped("output", "antipub")
		defer f.Close()
		for out := range outputCh {
			f.WriteString(out + "\r\n")
		}
	}()

	var u string
	for ap.Index, u = range urls {
		select {
		case <-utils.Done:
			goto off
		default:
		}
		utils.RateCounter.Incr(1)
		if ap.CheckPublic(u) {
			ap.Public++
			if viper.GetBool("antipub.savepublic") {
				outputCh <- u
			}
			continue
		}
		ap.Private++

		outputCh <- u
	}
	off:
	close(outputCh)
}