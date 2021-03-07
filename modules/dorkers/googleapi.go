package dorkers

import (
//	"context"
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	cse "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
//	"google.golang.org/api/option"
	"net/http"
	"net/url"
	"strings"
)

var Key = "AIzaSyCkq0momttTTPL5tUJZv8kNEGE3SXC_gYE"

func googleApiParse(dork string, page int) []string {
/*	bases := []string{
		"https://cse.google.com/cse/element/v1",
		//"https://www.google.co.uk/search",
		//"https://www.google.ru/search",
		//"https://www.google.fr/search",
	}*/

	//var Searcher *cse.CseService
	var out []string

	//ctrans := manager.ProxyToTransport(manager.PManager.ProxyFunc())

	apiTrans := &transport.APIKey{
		Key:       Key,
	//	Transport: ctrans,
	}

	client := http.DefaultClient
	client.Transport = apiTrans

	//if k, err := cse.NewService(context.Background(), option.WithHTTPClient(client),option.WithAPIKey("AIzaSyCkq0momttTTPL5tUJZv8kNEGE3SXC_gYE")); err == nil {

		//Searcher = cse.NewCseService(k)
	//}

	//var err error
	var res *cse.Search

	//if res, err = Searcher.List(dork).Num(100).Hl("en").Do(); err != nil {
	//	fmt.Println(err.Error())
	//	utils.ErrorCounter++
	//	return nil
	//}

	utils.RequestCounter++

	for _, item := range res.Items {
		out = append(out, item.Link)
	}
	utils.RateCounter.Incr(1)

/*

	for _, base := range bases {
		resp, err := manager.PManager.Get(buildGoogleApiUrl(dork, base, "en", page), uarand.GetRandom())
		if err != nil {
			continue
		}

		raw, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		buf := strings.TrimPrefix(strings.TrimSuffix(strings.ReplaceAll(strings.ReplaceAll(string(raw), "/*O_o/", ""), "\n", ""), ");"), "google.search.cse.api4846")

		var SearchResult *Search

		if err = json.Unmarshal([]byte(buf), &SearchResult); err == nil {
			for _, res := range SearchResult.Results {
				out = append(out, res.UnescapedURL)
			}
		}
	}*/
	return out
}

func buildGoogleApiUrl(searchTerm string, base string, languageCode string, index int) string {
	u, err := url.Parse(base)
	if err != nil {
		return ""
	}
	qu := u.Query()

	searchTerm = strings.Trim(searchTerm, " ")
	if index > 0 {
		index *= 100
	}
	// safe=off&cse_tok=AJvRUv1qwrWUX2B-AzQMFSCmUuro:1596075601195&sort=&exp=csqr,cc&oq=index.php%3Fid%3D1&gs_l=partner-generic.12...2144.4842.0.8333.14.14.0.0.0.0.144.1228.9j5.14.0.csems%2Cnrl%3D13...0.2664j974570j15...1.34.partner-generic..14.0.0.QYb9DNyIOO4&

	qu.Add("rsz", "filtered_cse")
	qu.Add("source", "gcsc")
	qu.Add("gss", ".com")
	qu.Add("cselibv", "26b8d00a7c7a0812")
	qu.Add("cx", "011463867123624065806:0aqhivzq0qu")
	qu.Add("q", searchTerm)
	qu.Add("oq", searchTerm)
	qu.Add("sort", "")
	qu.Add("exp", "csqr,cc")
	qu.Add("safe", "off")
	qu.Add("cse_tok", "AJvRUv1qwrWUX2B-AzQMFSCmUuro:1596075601195")
	qu.Add("num", fmt.Sprintf("%d", 100))
	qu.Add("hl", languageCode)
	qu.Add("start", fmt.Sprintf("%d", index))
	qu.Add("callback", "google.search.cse.api4846")

	u.RawQuery = qu.Encode()

	return u.String()
}