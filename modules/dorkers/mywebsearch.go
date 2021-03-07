package dorkers

import (
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"net/url"
	"strings"
)

/*
uri.addQueryParameter("p2","^MYWEBSEARCHDEFAULT^^^");
    uri.addQueryParameter("ln", "en");
    uri.addQueryParameter("tpr", "hpsb");
    uri.addQueryParameter("trs", "wtt");
    uri.addQueryParameter("searchfor",dork);
    uri.addQueryParameter("st", "hp");

 */

func mwsParse(dork string, page int, ret *chan []string) {
	base := "https://search.mywebsearch.com/mywebsearch/GGmain.jhtml"
	u, err := url.Parse(base)
	if err != nil {
		return
	}
	qu := u.Query()

	qu.Add("searchfor", dork)
	qu.Add("p2", "^MYWEBSEARCHDEFAULT^^^")
	qu.Add("ln", "en")
	qu.Add("tpr", "hpsb")
	qu.Add("trs", "wtt")
	qu.Add("st", "hp")

	u.RawQuery = qu.Encode()

	for i:=0; i < page; i++ {
		resp, body, err := manager.PManager.Get(u.String(), uarand.GetRandom())
		if err != nil {
			return
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			return
		}

		var o []string

		doc.Find("a.algo-title").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if !strings.Contains(link, resp.Request.URL.Hostname()) && !strings.Contains(link, ".gov/") {
				o = append(o, strings.Trim(link, " "))
			}
		})

		*ret <- o

		if link, nextFound := doc.Find("div.pagination-next a").Attr("href"); nextFound {
			u, _ = u.Parse(link)
		} else {
			return
		}
	}
}