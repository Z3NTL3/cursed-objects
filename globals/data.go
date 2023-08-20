package globals

var (
	// Table   = make(map[string][]string, 0)
	UAS     = make([]string, 0)
	ACCEPTS = make([]string, 0)
	REFS    = make([]string, 0)
	HTTP	= make([]string, 0)
	HTTPS	= make([]string, 0)
	SOCKS4	= make([]string, 0)
	SOCKS5	= make([]string, 0)
	PROXIES = make([]Proxy, 0)
)
