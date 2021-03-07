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
uri.addQueryParameter("s_chn", "prt_bon");
		    uri.addQueryParameter("q", dork);
		    uri.addQueryParameter("nojs", "1");
		    uri.addQueryParameter("b", Poco::NumberFormatter::format(page * 100));
		    uri.addQueryParameter("pz", "100");
		    uri.addQueryParameter("bct", "0");
		    uri.addQueryParameter("xargs", "0");
		    uri.addQueryParameter("v_t", "na");
 */

func aolParse(dork string, page int, ret *chan []string) {
	base := "https://search.aol.com/aol/search"
	u, err := url.Parse(base)
	if err != nil {
		return
	}
	qu := u.Query()

	qu.Add("q", dork)
	qu.Add("s_chn", "prt_bon")
	qu.Add("nojs", "1")
	qu.Add("b", fmt.Sprintf("%d", page * 100))
	qu.Add("pz", "100")
	qu.Add("bct", "0")
	qu.Add("xargs", "0")
	qu.Add("v_t", "na")

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

		doc.Find("ol li div").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Find("a").Attr("href")
			if !strings.Contains(link, resp.Request.URL.Hostname()) {
				o = append(o, strings.Trim(link, " "))
			}
		})

		*ret <- o

		if link, foundNext := doc.Find("a.next").Attr("href"); foundNext {
			u, _ = u.Parse(link)
		} else {
			return
		}
	}
}
