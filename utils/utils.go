package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/SilverCory/discordrpc-go"
	"github.com/denisbrodbeck/machineid"
	. "github.com/logrusorgru/aurora"
	"github.com/valyala/bytebufferpool"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var (
	clear     map[string]func()
	debug *os.File
	GoodServer string
	DiscordRP *discordrpc.RPCConnection
)

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	DiscordRP = discordrpc.NewRPCConnection("781973381796069416")
	debug, _ = os.Create("debug.txt")
}

func LogInfo(msg string) {
	fmt.Printf("\t[%s] [%s] ~> %s\r\n", Magenta(time.Now().Format("15:04:05")), Cyan("II"), msg)
}
func LogWarn(msg string) {
	fmt.Printf("\t[%s] [%s] ~> %s\r\n", Magenta(time.Now().Format("15:04:05")), Yellow("!!"), BrightYellow(msg))
}
func LogError(msg string) {
	fmt.Printf("\t[%s] [%s] ~> %s\r\n", Magenta(time.Now().Format("15:04:05")), Red("EE"), BrightRed(msg))
}
func LogDebug(msg string) {
	//if runtime.GOOS == "linux" {
	debug.WriteString(msg + "\r\n==================================\r\n")
	//}
}

func CreateFileTimeStamped(name, dir string) *os.File {
	_ = os.Mkdir(dir, 0700)
	f, err := os.Create(fmt.Sprintf("%s/%s-%s.txt", dir, name, strings.ReplaceAll(time.Now().Format("2006-01-01 12:04:05PM"), ":", "-")))
	if err != nil {
		LogError(err.Error())
		return nil
	}
	return f
}

func GetHWID() string {
	hashID, err := machineid.ProtectedID("XDumpGO")
	//hostStat, err := host.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println("HWID",hashID)
	return hashID
}

func CleanOutput(data string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(data, "\r", ""), "\n", ""))
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func ReadResponse(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	var err error
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func StrInArr(haystack []string, needle string) bool {
	for _, k := range haystack {
		if k == needle {
			return true
		}
	}
	return false
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		r := int(rand.Int63n(int64(len(charset))))
		if r < 0 || r > len(charset) {
			r = 0
		}
		b[i] = charset[r]
	}
	return string(b)
}

func ArrHasBlank(arr []string) bool {
	for _, k := range arr {
		if len(k) == 0 {
			return true
		}
	}
	return false
}

func StringOfLength(length int) string {
	return StringWithCharset(length, charset)
}

func HexStr(str string) string {
	return hex.EncodeToString([]byte(str))
}

func GetStringInBetweenZ(str string, start string, end string) (result string) {
	var s int
	var e int
	s = strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e = strings.Index(str[s:], end)
	if e == -1 {
		return
	}

	return strings.Trim(str[s : s+e], " ")
}

func GetStringInBetween(str, start, end string) (result string) {
	sIndex := strings.Index(str, start)
	if sIndex == -1 {
		return ""
	}
	parts := strings.Split(str[sIndex:], start)
	for _, part := range parts {
		if strings.Contains(part, end) {
			return strings.Split(part, end)[0]
		}
	}
	return ""
}

func ExtractData(str string, sel string) []string {
	re := regexp.MustCompile(fmt.Sprintf("(?s)%s(.*?)%s", sel, sel))
	return re.FindAllString(str, -1)
}

func Extract(str, left, right string, maxlen int) string {
	re := regexp.MustCompile(fmt.Sprintf("(?s)%s(.*?)%s", left, right))
	for {
		str := re.FindString(str)
		if len(str) == 0 {
			return ""
		}
		if len(str) <= maxlen {
			return str
		}
	}
}

func FindXPathChunk(str, sel string) int {
	re := regexp.MustCompile(fmt.Sprintf("(?s)%s(3+)", sel))
	if matches := re.FindStringSubmatch(str); len(matches) > 0 {
		z := len(matches[0]) - ((len(sel) * 2) + 1)
		return z
	}
	return 54
}

func ExtractXPathChunk(str, sel string) string {
	re := regexp.MustCompile(fmt.Sprintf("(?s)%s(.*?)(%s)", sel, sel))
	matches := re.FindStringSubmatch(str)
	out := matches[0]
	out = strings.TrimSuffix(out, sel[:1])
	out = strings.TrimSuffix(out, sel[:2])
	out = strings.TrimSuffix(out, sel[:3])
	out = strings.TrimSuffix(out, sel[:4])

	return out
}

func GetInputStr() string {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := reader.ReadString('\n')
	return strings.ReplaceAll(strings.ReplaceAll(opt, "\r", ""), "\n", "")
}

func HasAny(thing string, shit []string) bool {
	for _, v := range shit {
		if strings.Contains(thing, v) {
			return true
		}
	}
	return false
}

func ArrContains(arr []string, needle string) bool {
	for _, k := range arr {
		if k == needle {
			return true
		}
	}
	return false
}

func AskYesNo(msg string) bool {
	fmt.Println("\t" + msg)
	a := GetInputStr()
	if HasAny(a, []string{"y", "Y", "Yes", "yes"}) {
		return true
	}
	return false
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func FileLines(path string) int {

	length := 0

	file, err := os.Open(path)

	if err != nil {
		file, _ = os.Create(path)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) != "" {
			length++
		}
	}

	return length
}

func SliceToFile(slice []string, path string) {

	file, _ := os.Create(path)

	for _, str := range slice {
		file.WriteString(str + "\n")
	}

}

func FileToSlice(path string, sanitize bool) []string {

	file, err := os.Open(path)

	if err != nil {
		file, _ = os.Create(path)
	}

	scanner := bufio.NewScanner(file)

	var slice []string

	if sanitize == true {

		//Remove blank spaces
		var duplicates []string
		for scanner.Scan() {
			if strings.TrimSpace(scanner.Text()) != "" {
				duplicates = append(duplicates, strings.TrimSpace(scanner.Text()))
			}
		}

		//Remove Duplicates
		var lastStr string
		for _, duplicate := range duplicates {
			if lastStr != duplicate {
				lastStr = duplicate
				slice = append(slice, duplicate)
			}
		}

		file.Close()
		return slice
	}
	file.Close()
	return slice
}

func LoadInfo() []int {

	paths := []string{"params/keywords.txt", "params/pformats.txt", "params/ptypes.txt", "params/searchfuncs.txt", "params/domains.txt", "patterns.txt"}

	var values []int

	for _, path := range paths {

		switch path {
		case "params/keywords.txt":
			values = append(values, len(FileToSlice(path, true)))
		case "params/pformats.txt":
			values = append(values, len(FileToSlice(path, true)))
		case "params/ptypes.txt":
			values = append(values, len(FileToSlice(path, true)))
		case "params/searchfuncs.txt":
			values = append(values, len(FileToSlice(path, true)))
		case "params/domains.txt":
			values = append(values, len(FileToSlice(path, true)))
		case "patterns.txt":
			values = append(values, len(FileToSlice(path, true)))
		default:
			LogError("DID NOT LOAD ANYTHING")
		}
	}

	return values
}

func CompareTwoStrings(stringOne, stringTwo string) float32 {
	removeSpaces(&stringOne, &stringTwo)

	if value := returnEarlyIfPossible(stringOne, stringTwo); value >= 0 {
		return value
	}

	firstBigrams := make(map[string]int)
	for i := 0; i < len(stringOne)-1; i++ {
		a := fmt.Sprintf("%c", stringOne[i])
		b := fmt.Sprintf("%c", stringOne[i+1])

		bigram := a + b

		var count int

		if value, ok := firstBigrams[bigram]; ok {
			count = value + 1
		} else {
			count = 1
		}

		firstBigrams[bigram] = count
	}

	var intersectionSize float32
	intersectionSize = 0

	for i := 0; i < len(stringTwo)-1; i++ {
		a := fmt.Sprintf("%c", stringTwo[i])
		b := fmt.Sprintf("%c", stringTwo[i+1])

		bigram := a + b

		var count int

		if value, ok := firstBigrams[bigram]; ok {
			count = value
		} else {
			count = 0
		}

		if count > 0 {
			firstBigrams[bigram] = count - 1
			intersectionSize = intersectionSize + 1
		}
	}

	return (2.0 * intersectionSize) / (float32(len(stringOne)) + float32(len(stringTwo)) - 2)
}

func removeSpaces(stringOne, stringTwo *string) {
	*stringOne = strings.Replace(*stringOne, " ", "", -1)
	*stringTwo = strings.Replace(*stringTwo, " ", "", -1)
}

func returnEarlyIfPossible(stringOne, stringTwo string) float32 {
	// if both are empty strings
	if len(stringOne) == 0 && len(stringTwo) == 0 {
		return 1
	}

	// if only one is empty string
	if len(stringOne) == 0 || len(stringTwo) == 0 {
		return 0
	}

	// identical
	if stringOne == stringTwo {
		return 1
	}

	// both are 1-letter strings
	if len(stringOne) == 1 && len(stringTwo) == 1 {
		return 0
	}

	// if either is a 1-letter string
	if len(stringOne) < 2 || len(stringTwo) < 2 {
		return 0
	}

	return -1
}

func SplitSlice(slice []string, n int) [][]string {
	var divided [][]string
	if n == 1 {
		return append(divided, slice)
	}

	chunkSize := (len(slice) + n - 1) / n

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		divided = append(divided, slice[i:end])
	}
	return divided
}

func StatusStr(status int) string {
	switch status {
	case 0:
		return "Running"
	case 1:
		return "Stopping"
	}
	return "Unknown"
}

func DoCleanUrls(input []string) []string {
	var hosts []string
	var out []string
	for _, str := range input {
		//if m, err := regexp.MatchString(`(.*?)\?.+=`, str); err != nil && m {
		if !(strings.Contains(str, "http") && strings.Contains(str, "=") && strings.Contains(str, "?")) {
			continue
		}
		u, err := url.Parse(str)
		if err != nil {
			continue
		}
		if !StrInArr(hosts, u.Host) {
			hosts = append(hosts, u.Host)
			out = append(out, strings.ReplaceAll(str, "&amp;", "&"))
		}
		//}
	}
	return out
}

func GetFileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes), nil
}

func IsClosed(ch <-chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}