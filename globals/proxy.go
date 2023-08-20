package globals

type ProtocolType string

const (
	_HTTP   ProtocolType = "http"
	_HTTPS  ProtocolType = "https"
	_SOCKS4 ProtocolType = "socks4"
	_SOCKS5 ProtocolType = "socks5"
)

type Proxy struct {
	Protocol ProtocolType
	ProxyStr string
}
