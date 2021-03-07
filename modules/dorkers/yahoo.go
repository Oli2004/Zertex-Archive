package dorkers

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"net/http"
	"net/url"
	"strings"
)

func yahooParse(dork string, page int, ret *chan []string) {
	base := "https://search.yahoo.com/search"

	u, err := url.Parse(base)
	if err != nil {
		return
	}
	qu := u.Query()

	// p=index.php&fr=yfp-t&fp=1&toggle=1&cop=mss&ei=UTF-8
	// https://search.yahoo.com/search?p=index.php&ei=UTF-8&fr=yfp-t&fp=1&b=11&pz=100&bct=0&xargs=0

	qu.Add("p", dork)
	qu.Add("fr", "yfp-t")
	qu.Add("fp","1")
	qu.Add("toggle","1")
	qu.Add("cop","mss")
	qu.Add("ei","UTF-8")
	qu.Add("pz", "100")
	qu.Add("xargs", "0")
	qu.Add("b", fmt.Sprintf("%d", (page * 100) + 1))
	qu.Add("bct", "0")

	u.RawQuery = qu.Encode()


	headers := http.Header{}
	headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	headers.Add("Accept-Language", "en-US,en;q=0.5")
	headers.Add("Referrer", "https://www.yahoo.com/")

	for i:=0; i < page; i ++ {
		resp, body, err := manager.PManager.GetWithHeaders(u.String(), uarand.GetRandom(), headers)
		if err != nil {
			return
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			return
		}

		var o []string

		doc.Find("h3.title.ov-h").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Find("a").Attr("href")
			if !strings.Contains(link, resp.Request.URL.Hostname()) && !strings.Contains(link, ".gov/") {
				o = append(o, strings.Trim(link, " "))
			}
		})

		*ret <- o

		if link, found := doc.Find("a.next").Attr("href"); found {
			u, _ = u.Parse(link)
		} else {
			return
		}
	}
}
