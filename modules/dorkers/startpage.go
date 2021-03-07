package dorkers

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func startpageParse(dork string, page int, ret *chan []string)  {
	base := "https://www.startpage.com/sp/search"

	for i := 0; i < page; i++ {
		resp, err := manager.PManager.Post(base, "application/x-www-form-urlencoded", fmt.Sprintf("query=%s&language=english&lui=english&abp=-1&cat=web&t=default&page=%d", dork, i+1))
		if err != nil {
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return
		}

		var o []string
		doc.Find("a.w-gl__result-title").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if !strings.Contains(link, resp.Request.URL.Hostname()) {
				o = append(o, strings.Trim(link, " "))
			}
		})

		*ret <- o
	}
}