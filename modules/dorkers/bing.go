package dorkers

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

func bingParse(dork string, page int, ret *chan []string) {
	base := "https://www.bing.com/search"

	u, err := url.Parse(base)
	if err != nil {
		return
	}
	qu := u.Query()

	qu.Add("q", dork)
	qu.Add("count", "50")
	qu.Add("qs", "ds")
	qu.Add("go", "Search")
	qu.Add("form", "QBRE")
	//qu.Add("count", "100")
	//qu.Add("first", fmt.Sprintf("%d", page * 15))
	u.RawQuery = qu.Encode()

	headers := http.Header{}
	headers.Add("SRCHHPGUSR", "NEWWND=0&NRSLT=50&SRCHLANG=&AS=1&ADLT=DEMOTE&NNT=1&BRW=W&BRH=S&CW=1920&CH=560&DPR=1&UTC=-300&DM=0&HV=1604285262&WTS=63739882180")

	for i:=0; i < page; i++ {
		_, body, err := manager.PManager.GetWithHeaders(u.String(), "Nokia2700c/10.0.011 (SymbianOS/9.4; U; Series60/5.0 Opera/5.0; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Safari/525 3gpp-gba", headers)
		if err != nil {
			return
		}

		/*if utils.HasAny(body, []string{
			"CHGEDGE1708",
		}) {
			goto er
		}*/

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			return
		}

		var o []string

		sel := doc.Find("li.b_algo")
		for i := range sel.Nodes {
			item := sel.Eq(i)
			linkTag := item.Find("a")
			link, _ := linkTag.Attr("href")
			link = strings.Trim(link, " ")
			if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
				o = append(o, link)
			}
		}

		*ret <- o

		qu.Set("start", fmt.Sprintf("%d", (i*10) + 1))
		/*if link, foundNext := doc.Find("a.sb_pagN").Attr("href"); foundNext {
			u, _ = u.Parse(link)
		} else {
			fmt.Println(dork,"ended on",i)
			return
		}*/
	}
}