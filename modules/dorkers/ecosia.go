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
uri.addQueryParameter("q", dork);
    uri.addQueryParameter("p", Poco::NumberFormatter::format(page));
    uri.addQueryParameter("c", "en");

 */

func ecosiaParse(dork string, page int, ret *chan []string) {
	base := "https://www.ecosia.org/search"

	u, err := url.Parse(base)
	if err != nil {
		return
	}
	qu := u.Query()

	qu.Add("q", dork)
	qu.Add("p", fmt.Sprintf("%d", page))
	qu.Add("c", "en")

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

		doc.Find("a.result-snippet-link").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if !strings.Contains(link, resp.Request.URL.Hostname()) {
				o = append(o, strings.Trim(link, " "))
			}
		})

		*ret <- o

		if link, found := doc.Find("a.pagination-button.pagination-next").Attr("href"); found {
			u.Parse(link)
		} else {
			return
		}
	}
}