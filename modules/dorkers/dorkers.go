package dorkers

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type Dorker struct {
	Name string
	Enabled bool
	Scrape func(dork string, page int, ret *chan []string)
	parse func(resp *http.Response) []string
	Bases []string
	Return chan []string
}

var Dorkers map[string]*Dorker

func SetupDorkers() {
	Dorkers = make(map[string]*Dorker)
	Dorkers["Google"] = NewDorker("Google", googleParse) // done
	//Dorkers["GoogleAPI"] = NewDorker("GoogleAPI", googleApiParse)
	Dorkers["Bing"] = NewDorker("Bing", bingParse) // done
	Dorkers["AOL"] = NewDorker("AOL", aolParse) // done
	Dorkers["MyWebSearch"] = NewDorker("MyWebSearch", mwsParse) // done
	Dorkers["Yahoo"] = NewDorker("Yahoo", yahooParse) // done
	Dorkers["DuckDuckGo"] = NewDorker("DuckDuckGo", ddgParse)
	Dorkers["Ecosia"] = NewDorker("Ecosia", ecosiaParse)
	Dorkers["Qwant"] = NewDorker("Qwant", qwantParse) // eh?
	Dorkers["StartPage"] = NewDorker("StartPage", startpageParse) // done
	Dorkers["Yandex"] = NewDorker("Yandex", yandexParse) // done

	for name, k := range Dorkers {
		k.Enabled = viper.GetBool(fmt.Sprintf("scraper.Engines.%s", name))
	}
}

func NewDorker(name string, Func func(dork string, page int, ret *chan []string)) *Dorker {
	return &Dorker{
		Name: name,
		Scrape: Func,
		Return: make(chan []string),
	}
}