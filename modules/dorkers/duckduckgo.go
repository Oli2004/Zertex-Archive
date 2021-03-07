package dorkers

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

/*
uri.addQueryParameter("s", Poco::NumberFormatter::format((page * 50) + 30));
    uri.addQueryParameter("dc", Poco::NumberFormatter::format((page * 30) + 1));
    uri.addQueryParameter("o", "json");
    uri.addQueryParameter("api", "/d.js");
    uri.addQueryParameter("v", "l");
    uri.addQueryParameter("kl", "wt-wt");

 */

func ddgParse(dork string, page int, ret *chan []string) {
	base := "https://html.duckduckgo.com/html"

	u, err := url.Parse(base)
	if err != nil {
		return
	}
	qu := u.Query()

	qu.Add("s", fmt.Sprintf("%d", (page * 50) + 30))
	qu.Add("dc", fmt.Sprintf("%d", (page * 30) + 1))
	qu.Add("o", "json")
	qu.Add("api", "/d.js")
	qu.Add("v", "l")
	qu.Add("kl", "wt-wt")

	u.RawQuery = qu.Encode()

	for i:=0; i < page; i++ {
		resp, err := manager.PManager.Post(u.String(), "application/x-www-form-urlencoded", fmt.Sprintf("q=%s&b=&kl=&df=", dork))
		if err != nil {
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return
		}

		var o []string

		doc.Find("h2.result__title").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Find("a").Attr("href")
			if !strings.Contains(link, resp.Request.URL.Hostname()) {
				o = append(o, strings.Trim(link, " "))
			}
		})

		*ret <- o
	}
}