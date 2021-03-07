package manager

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/api/client"
	"git.quartzinc.dev/Zertex/XDGv2/qtui"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	"github.com/corpix/uarand"
	"github.com/gosuri/uilive"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const
(
	HTTP = iota
	SOCKS4
	SOCKS5
)

type Proxy struct {
	Address string
	Type int
	Alive bool
	Sem chan interface{}
	Auth string
	Transport *http.Transport
}

func (p *Proxy) ToProxyUrl() (*url.URL, error) {
	var scheme string
	switch p.Type {
	case HTTP:
		scheme = "http"
		break
	case SOCKS4:
		scheme = "socks4"
		break
	case SOCKS5:
		scheme = "socks5"
		break
	}

	_u, err := url.Parse(fmt.Sprintf("%s://%s", scheme, p.Address))
	if err != nil {
		return nil, err
	}

	if len(p.Auth) > 0 {
		_u.User = url.UserPassword(strings.Split(p.Auth, ":")[0], strings.Split(p.Auth, ":")[1])
	}
	return _u, nil
}

func (p *Proxy) ToUrlFunc() func(r *http.Request) (*url.URL, error) {
	return func(r *http.Request) (*url.URL, error) {
		return p.ToProxyUrl()
	}
}

type ProxyManager struct {
	Proxies []*Proxy
	ctx context.Context
	cancel context.CancelFunc
	X chan interface{}
	Client *http.Client
	torConn *net.TCPConn

	FastClient *fasthttp.Client
}

func (pm *ProxyManager) RotateTorIP() {
	pm.torConn.Write([]byte("authenticate '\"\"'\n"))
	pm.torConn.Write([]byte("signal newnym\n"))
	pm.torConn.Write([]byte("quit\n"))
}

var torNewIP = "echo authenticate '\"\"'; echo signal newnym; echo quit"

type Req struct {
	Request *http.Request
	Host string
	Response *chan *http.Response
}

var PManager *ProxyManager

func init() {
	ctx, cancel := context.WithCancel(context.Background())

	PManager = &ProxyManager{
		Proxies: make([]*Proxy, 0),
		ctx: ctx,
		cancel: cancel,
		Client: &http.Client{/*Timeout: time.Duration(viper.GetInt("core.Timeouts")) * time.Second,*/ Jar: nil},
			FastClient: &fasthttp.Client{
			MaxConnsPerHost:               0,
			MaxIdleConnDuration:           -1,
			MaxResponseBodySize:           2e8,
			RetryIf: func(request *fasthttp.Request) bool {
				return false
			},
		},
	}
	PManager.Client.Transport = PManager.CreateProxyTransport()
}

func (pm *ProxyManager) ProxyStringArray() []string {
	var list []string
	for _, proxy := range pm.Proxies {
		if len(proxy.Auth) > 0 {
			list = append(list, fmt.Sprintf("%s:%s", proxy.Address, proxy.Auth))
		} else {
			list = append(list, proxy.Address)
		}
	}
	return list
}

func (pm *ProxyManager) LoadScanner(scanner *bufio.Scanner, proxyType int) int {
		for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		var p *Proxy
		if len(parts) == 4 {
			p = &Proxy{Address: fmt.Sprintf("%s:%s", parts[0], parts[1]), Auth: fmt.Sprintf("%s:%s", parts[2], parts[3]), Type: proxyType, Alive: true, Sem: make(chan interface{}, 10)}
		} else {
			p = &Proxy{Address: line, Type: proxyType, Alive: true, Sem: make(chan interface{}, 10)}
		}
		p.Transport = &http.Transport{
			TLSClientConfig:        &tls.Config{
				InsecureSkipVerify: true,
				VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
					return nil
				},
			},
			DialContext: (&net.Dialer{
				Timeout: 10*time.Second,
			}).DialContext,
			DialTLSContext: (&net.Dialer{
				Timeout: 10*time.Second,
			}).DialContext,
			Proxy:                  p.ToUrlFunc(),
			TLSHandshakeTimeout:    10 * time.Second,
			DisableCompression:     false,
			ForceAttemptHTTP2:      false,
			MaxIdleConns:           100,
			MaxIdleConnsPerHost:    35,
			MaxConnsPerHost:        35,
			IdleConnTimeout:        10 * time.Second,
			ResponseHeaderTimeout:  time.Duration(viper.GetInt("core.timeout")) * time.Second,
			TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
		}
		pm.Proxies = append(pm.Proxies, p)
	}
	return len(pm.Proxies)
}

func (pm *ProxyManager) LoadFile(filename string, proxyType int) int {
	file, err := os.Open(filename)
	if err != nil {
		utils.LogError(err.Error())
		return 0
	}
	defer file.Close()
	pm.Proxies = []*Proxy{}
	scanner := bufio.NewScanner(file)
	return pm.LoadScanner(scanner, proxyType)
}

func (pm *ProxyManager) LoadCustomerProxies(loading *qtui.LoadingWindow) int {
	loading.ProgressBar.SetMaximum(2)
	loading.ProgressBar.SetValue(0)
	loading.Label.SetText("Contacting server...")
	pm.Proxies = []*Proxy{}
	if client.XDGAPI.RequestProxies() != nil {
		loading.Label.SetText("Failed to contact server.")
		return -1
	}
	loading.ProgressBar.SetValue(1)
	loading.Label.SetText("Downloading list...")

	proxies := <- client.XDGAPI.Proxies()
	loading.ProgressBar.SetValue(2)

	loading.Label.SetText("Loading proxies...")
	loading.ProgressBar.SetMaximum(len(proxies.List))
	f := make(chan interface{})
	go func() {
		for {
			select {
			case <-time.After(100 * time.Millisecond):
				loading.ProgressBar.SetValue(len(pm.Proxies))
			case <-f:
				loading.ProgressBar.SetValue(len(pm.Proxies))
				return
			}
		}
	}()
	for _, line := range proxies.List {
		p := strings.Split(line, ":")
		pr := &Proxy{
			Address: strings.Join(p[:2], ":"),
			Type:    SOCKS5,
			Alive:   true,
			Sem:     make(chan interface{}, 10),
			Auth:    strings.Join(p[2:], ":"),
		}
		pm.Proxies = append(pm.Proxies, pr)
	}

	return len(pm.Proxies)
}

func (pm *ProxyManager) TestProxies() {
	// http://azenv.net/
	var living []*Proxy
	startTime := time.Now()
	dead := 0
	wg := sync.WaitGroup{}
	threads := viper.GetInt("core.Threads")
	sem := make(chan int, threads)
	proxyChan := make(chan *Proxy)
	done := make(chan interface{})
	kill := make(chan interface{})
	ctx, cancel := context.WithCancel(context.Background())
	defer close (sem)
	go func() {
		writer := uilive.New()
		writer.Start()
		defer writer.Stop()
		for len(sem) > 0 {
			select {
			case <-utils.CancelChan:
				close(done)
				for len(sem) > 0 {
					select {
					case <-utils.CancelChan:
						cancel()
						utils.LogError("Killing threads...")
						close(kill)
						return
					default:
						fmt.Fprintf(writer, "\t[%s] (%d) Waiting on threads to settle down ~> [Alive: %d] [Dead: %d]\r\n", Magenta(utils.FmtDuration(time.Now().Sub(startTime))), len(sem), len(living), dead)
						time.Sleep(250 * time.Millisecond)
					}
				}
				return
			default:
				fmt.Fprintf(writer, "\t[%s] [%d/%d] (%d) Testing... ~> [Alive: %d] [Dead: %d]\r\n", Magenta(utils.FmtDuration(time.Now().Sub(startTime))), dead+len(living), len(pm.Proxies), len(sem), len(living), dead)
				time.Sleep(250 * time.Millisecond)
			}
		}
	}()

	wg.Add(threads)

	for i:=0; i< threads; i++ {
		sem<-1
		go func() {
			defer func() {
				wg.Done()
				<-sem
			}()

			for proxy := range proxyChan {
				if proxy == nil {
					return
				}
				req, err := http.NewRequest("GET", "https://azenv.net", nil)
				if err != nil {
					return
				}

				reqctx := req.WithContext(ctx)
				//utils.LogInfo("testing: " + proxy.Address)
				httpClient := &http.Client{Transport: &http.Transport{
					Proxy: proxy.ToUrlFunc(),
				}}
				reqctx.Header.Add("User-Agent", uarand.GetRandom())
				resp, err := httpClient.Do(reqctx)
				if err != nil {
					dead++
					continue
				}
				resp.Body.Close()
				//LogInfo(fmt.Sprintf("status: %d", resp.StatusCode))
				if resp.StatusCode != http.StatusOK {
					dead++
				} else {
					proxy.Alive = true
					living = append(living, proxy)
				}
			}
		}()
	}

	for _, k := range pm.Proxies {
		select {
		case <- done:
			goto popoff
		case proxyChan <- k:
			continue
		}
	}
popoff:
	close(proxyChan)
	f := make(chan interface{})
	go func() {
		wg.Wait()
		close(f)
	}()
	select {
	case <- f:
		break
	case <- kill:
		break
	}

	var d string
	for _,k := range living {
		d += k.Auth + ":" + k.Address + "\r\n"
	}
	ioutil.WriteFile("proxies-living.txt", []byte(d), 0644)
	//pm.Proxies = living
	utils.LogInfo("Finished testing proxies")
	time.Sleep(2 * time.Second)
}

func (pm *ProxyManager) ProxyFunc(r *http.Request) (*url.URL, error) {
	if proxy, err := pm.GetRandomProxy(); err == nil {
		return proxy.ToProxyUrl()
	} else {
		return nil, nil
	}
}

func (pm *ProxyManager) GetRandomProxy() (*Proxy, error) {
	if len(pm.Proxies) == 0 {
		return nil, errors.New("no proxies")
	}
	var proxy *Proxy

	if len(pm.Proxies) == 1 {
		proxy = pm.Proxies[0]
	} else {
		proxy = pm.Proxies[rand.Intn(len(pm.Proxies)-1)]
	}
	return proxy, nil
}

func (pm *ProxyManager) CreateProxyTransport() *http.Transport {
	//customTransport := http.DefaultTransport.(*http.Transport).Clone()
	return &http.Transport{
		TLSClientConfig:        &tls.Config{
			InsecureSkipVerify: true,
			VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
				return nil
			},
		},
		DialContext: (&net.Dialer{
			Timeout: 5*time.Second,
			KeepAlive: -1,
		}).DialContext,
		Proxy:                  pm.ProxyFunc,
		TLSHandshakeTimeout:    5 * time.Second,
		DisableKeepAlives:      true,
		DisableCompression:     false,
		ForceAttemptHTTP2:      false,
		MaxIdleConns:           -1,
		MaxIdleConnsPerHost:    -1,
		//MaxConnsPerHost:        35,
		IdleConnTimeout:        time.Nanosecond,
		ExpectContinueTimeout:  4 * time.Second,
		ResponseHeaderTimeout:  time.Duration(viper.GetInt("core.timeout")) * time.Second,
		TLSNextProto:           make(map[string]func(string, *tls.Conn) http.RoundTripper),
	}
}

func (pm *ProxyManager) ResetCtx() {
	pm.ctx, pm.cancel = context.WithCancel(context.Background())
}

func (pm *ProxyManager) CancelAll() {
	if pm.ctx == nil {
		utils.LogError("ctx is nil")
		return
	}
	pm.cancel()
	pm.ResetCtx()
}

func (pm *ProxyManager) Post(url string, _type string, data string) (*http.Response, error) {
	timeout := viper.GetInt("core.Timeouts")
	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	httpClient := &http.Client{Transport: pm.CreateProxyTransport(), Timeout: time.Duration(timeout) * time.Second}
	for attempt := 0; attempt < 3; attempt++ {
		if pm.ctx == nil {
			break
		}
		reqctx := req.WithContext(pm.ctx)
		reqctx.Header.Add("Content-Type", _type)
		reqctx.Header.Add("User-Agent", uarand.GetRandom())
		//resp, err := httpClient.Post(url, _type, strings.NewReader(data))
		resp, err := httpClient.Do(reqctx)
		if err != nil {
			continue
		}

		return resp, nil
	}
	return nil, errors.New("failed post request")
}


func (pm *ProxyManager) Get(_url string, ua string) (*http.Response, string, error) {
	s := make(chan interface{})
	defer close(s)
	return pm.BaseGet(_url, ua, 0, http.Header{}, true, nil, &s)
}

func (pm *ProxyManager) GetWithErrors(_url string, ua string, remote *int) (*http.Response, string, error) {
	s := make(chan interface{})
	defer close(s)
	return pm.BaseGet(_url, ua, 0, http.Header{}, true, remote, &s)
}
func (pm *ProxyManager) GetWithErrorsSkip(_url string, ua string, remote *int, skip *chan interface{}) (*http.Response, string, error) {
	return pm.BaseGet(_url, ua, 0, http.Header{}, true, remote, skip)
}

func (pm *ProxyManager) GetWithoutFail(_url, ua string) (*http.Response, string, error) {
	s := make(chan interface{})
	defer close(s)
	return pm.BaseGet(_url, ua, 0, http.Header{}, false, nil, &s)
}

func (pm *ProxyManager) GetWithoutFailWithErrors(_url, ua string, remote *int, skip *chan interface{}) (*http.Response, string, error) {
	return pm.BaseGet(_url, ua, 0, http.Header{}, false, remote, skip)
}

func (pm *ProxyManager) GetWithTimeout(_url string, ua string, extratimeout int) (*http.Response, string, error) {
	s := make(chan interface{})
	defer close(s)
	return pm.BaseGet(_url, ua, extratimeout, http.Header{}, true, nil, &s)
}

func (pm *ProxyManager) GetWithHeaders(_url string, ua string, headers http.Header) (*http.Response, string, error) {
	s := make(chan interface{})
	defer close(s)
	return pm.BaseGet(_url, ua, 0, headers, true, nil, &s)
}

func (pm *ProxyManager) BaseFastGet(_url, ua string, headers http.Header, remoteError *int, canfail bool) (int, *fasthttp.ResponseHeader, string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(_url)

	req.Header.SetUserAgent(ua)
	for key, val := range headers {
		req.Header.Set(key, strings.Join(val, ";"))
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	proxy, err := pm.GetRandomProxy()
	if err != nil {
		return -1, nil, "", err
	}

	switch proxy.Type {
	case HTTP:
		pm.FastClient.Dial = fasthttpproxy.FasthttpHTTPDialer(fmt.Sprintf("%s@%s", proxy.Auth, proxy.Address))
	case SOCKS4:
		fallthrough
	case SOCKS5:
		pm.FastClient.Dial = fasthttpproxy.FasthttpSocksDialer(fmt.Sprintf("%s@%s", proxy.Auth, proxy.Address))
	}

	for i:=0; i < 3; {
		if utils.Module != "Dumper" && canfail {
			i++
		}
		err := pm.FastClient.DoTimeout(req, resp, time.Duration(viper.GetInt("core.timeouts"))*time.Second)
		if err != nil {
			utils.ErrorCounter++
			if utils.HasAny(err.Error(), []string {
				"socks connect",
				"read tcp",
			}) {
				i--
			}
			continue
		}
		utils.RequestCounter++
		utils.RateCounter.Incr(1)

		contentEncoding := resp.Header.Peek("Content-Encoding")
		var body []byte
		if bytes.EqualFold(contentEncoding, []byte("gzip")) {
			fmt.Println("Unzipping...")
			body, _ = resp.BodyGunzip()
		} else {
			body = resp.Body()
		}

		return resp.StatusCode(), &resp.Header, string([]byte(string(body))), nil
	}
	return -1, nil, "", errors.New("failed get")
}

func (pm *ProxyManager) BaseGet(_url string, ua string, extratimeout int, headers http.Header, canfail bool, remoteErrors *int, skip *chan interface{}) (*http.Response, string, error) {
	req, err := http.NewRequestWithContext(pm.ctx, "GET", _url, nil)
	if err != nil {
		return nil, "", err
	}

	if len(headers) > 0 {
		for key, vals := range headers {
			req.Header.Set(key, strings.Join(vals, ";"))
		}
	}
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Connection", "close")
	req.Close = true

	for attempt := 0; attempt < 3;  {
		if utils.Module != "Dumper" && canfail {
			attempt++
		}
		select {
		case <-utils.Done:
			return nil, "", errors.New("exit early")
		case <- *skip:
			return nil, "", errors.New("skipped")
		default:
			proxy, err := pm.GetRandomProxy()
			if err != nil {
				continue
			}
			resp, err := proxy.Transport.RoundTrip(req)
			if err != nil {
				if resp != nil {
					resp.Body.Close()
				}
				utils.ErrorCounter++
				if remoteErrors != nil {
					*remoteErrors++
				}
				continue
			}

			utils.RateCounter.Incr(1)
			utils.RequestCounter++
			res, err := utils.ReadResponse(resp)
			if err != nil {
				fmt.Println("err read resp", err.Error())
				continue
			}
			return resp, res, err
		}
	}
	return nil, "", errors.New("failed request")
}
