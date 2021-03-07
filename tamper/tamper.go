package tamper

import "strings"

var tampers map[string]func(query string) string

func init() {
	tampers = make(map[string]func(query string) string)
	tampers["appendnull"] = AppendNullByte
	tampers["base64"] = B64Encode
	tampers["bluecoat"] = Bluecoat
	tampers["equal2like"] = Equal2Like
	tampers["space2comment"] = Space2Comment
	tampers["space2mysqlblank"] = Space2MySQLBlank
	tampers["space2mysqldash"] = Space2MySQLDash
	tampers["apostraphenullencode"] = ApostropheNullEncode
	tampers["between"] = Between
	tampers["modsec"] = ModSec
	tampers["concat2concatws"] = Concat2ConcatWS
	tampers["luanginx"] = LuaNGINX
	tampers["randomcase"] = RandomCase
	tampers["lowercase"] = LowerCase
	tampers["uppercase"] = UpperCase
	tampers["booleanmask"] = BooleanMask
}

func Tamper(query string, ss []string) string {
	//fmt.Printf("%s tampering with %v\r\n", query, ss)
	for _,tamper := range ss {
		//fmt.Println("tamper", tamper)
		query = tampers[tamper](query)
	}
	return strings.Replace(query, "%00", "", -1)
}