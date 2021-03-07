package dorkers

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"net/url"
	"strings"
)

func yandexParse(dork string, page int, ret *chan []string) {
	base := "https://yandex.com/search/"

	u, err := url.Parse(base)
	if err != nil {
		return
	}
	qu := u.Query()

	qu.Add("text", dork)
	qu.Add("p", fmt.Sprintf("%d", page))

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

		doc.Find("a.link.link_theme_normal.organic__url.link_cropped_no.i-bem").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if !strings.Contains(link, resp.Request.URL.Hostname()) && !strings.Contains(link, ".gov/") {
				o = append(o, strings.Trim(link, " "))
			}
		})

		*ret <- o

		if link, found := doc.Find("a.link.link_theme_none.link_target_serp.pager__item.pager__item_kind_next.i-bem").Attr("href"); found {
			u, _ = u.Parse(link)
		} else {
			return
		}
	}
}