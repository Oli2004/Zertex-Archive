package dorkers

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"net/url"
	"strings"
)

/*
uri.addQueryParameter("count", "10");
    uri.addQueryParameter("offset", Poco::NumberFormatter::format(page * 10));
    uri.addQueryParameter("q", dork);
    uri.addQueryParameter("t", "web");
    uri.addQueryParameter("device", "desktop");
    uri.addQueryParameter("safesearch", "0");
    uri.addQueryParameter("locale", "en_US");
    uri.addQueryParameter("uiv", "4");

 */

func qwantParse(dork string, page int, ret *chan []string) {
	defer close(*ret)
	base := "https://lite.qwant.com/"
	u, err := url.Parse(base)
	if err != nil {
		return
	}
	qu := u.Query()

	qu.Add("q", dork)
	qu.Add("count", "10")
	qu.Add("offset", fmt.Sprintf("%d", page * 10))
	qu.Add("t", "web")
	qu.Add("device", "desktop")
	qu.Add("safesearch", "0")
	qu.Add("locale", "en_US")
	qu.Add("uiv", "4")

	u.RawQuery = qu.Encode()

	for i := 0; i < page; i++ {

		resp, body, err := manager.PManager.Get(u.String(), uarand.GetRandom())
		if err != nil {
			return
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			return
		}

		var o []string

		doc.Find("div.url").Each(func(i int, s *goquery.Selection) {
			link := s.Text()
			link = strings.ReplaceAll(strings.ReplaceAll(link, "<b>", ""), "</b>", "")
			if !strings.Contains(link, resp.Request.URL.Hostname()) && !strings.Contains(link, ".gov/") {
				o = append(o, strings.Trim(link, " "))
			}
		})
		*ret <- o

	}
}