package modules

import (
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"git.quartzinc.dev/Zertex/XDGv2/modules/dorkers"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	"github.com/paulbellamy/ratecounter"
	"github.com/spf13/viper"
	"net/url"
	"sync"
	"time"
)

type ScrapeModule struct {
	Index int
	Dorks []string
	Urls []string
	UrlCh chan []string
	DorkCh chan string
}

var Scraper *ScrapeModule

func (sm *ScrapeModule) Start() {
	manager.PManager.ResetCtx()
	utils.RequestCounter = 0
	utils.ErrorCounter = 0
	utils.RateCounter = ratecounter.NewRateCounter(1 * time.Second)
	utils.StartTime = time.Now()
	threads := viper.GetInt("core.Threads")
	utils.GlobalSem = make(chan interface{}, threads)
	pages := viper.GetInt("scraper.Pages")

	sm.UrlCh = make(chan []string, threads)
	sm.DorkCh = make(chan string)
	utils.Done = make(chan interface{})
	utils.Kill = make(chan interface{})

	go func() {
		f := utils.CreateFileTimeStamped("urls", "urls")
		for urls := range sm.UrlCh {
			for _, _u := range urls {
				if _u, err := url.Parse(_u); err == nil {
					if viper.GetBool("scraper.Filter") {
						if len(_u.RawQuery) == 0 ||
							utils.HasAny(_u.Hostname(), []string{"facebook", "google.com", "stackoverflow.com", "php.net", "yahoo.com", "bing.com", ".gov", "youtube.com", "yandex.com"}) ||
							len(_u.Query()) == 0 {
							continue
						}
					}
				} else {
					continue
				}
				sm.Urls = append(sm.Urls, _u)
				f.WriteString(_u + "\r\n")
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(threads)
	for i:=0; i < threads; i++ {
		utils.GlobalSem<-0
		go func() {
			defer func() {
				wg.Done()
				<-utils.GlobalSem
			}()
			for dork := range sm.DorkCh {
				for _, dorker := range dorkers.Dorkers {
					if dorker.Enabled {
						dorker.Scrape(dork, pages, &sm.UrlCh)
					}
				}
			}
		}()
	}

	var dork string
	for sm.Index, dork = range sm.Dorks {
		select {
		case <-utils.Done:
			goto popoff
		case sm.DorkCh <- dork:
		}
	}
popoff:
	close(sm.DorkCh)
	f := make(chan interface{})
	go func() {
		wg.Wait()
		close(f)
	}()
	select {
	case <- f:
		sm.Index = 0
		break
	case <- utils.Kill:
		break
	}
}