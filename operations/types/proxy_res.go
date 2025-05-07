package types

// ProxyInfo represents proxy information.
type ProxyInfo struct {
	ID                int    `json:"id"`                // Primary key
	ExpiryTime        int    `json:"expiryTime"`        // Expiration time (timestamp)
	ProxyCategoryType int    `json:"proxyCategoryType"` // Proxy type: 1: Platform proxy; 2: Self-owned proxy
	ProxyCheckStatus  int    `json:"proxyCheckStatus"`  // Detection status: 0-pending detection 1-monitoring success 2-detection failure 3-unknown error
	ProxyIP           string `json:"proxyIp"`           // Proxy IP
	ProxyName         string `json:"proxyName"`         // Proxy Name
	ProxyProvider     int    `json:"proxyProvider"`     // Proxy Provider: default value 0 - none 4-Oxylabs 5-Proxys.io 7-Luminati 8-Lumauto 9-Oxylabsauto 10-Trojan，11-Shadowsocks 13-ABCPROXY 14-LunaProxy 15-IPHTML 16-PiaProxy 17-922S5 18-360Proxy
	ProxyType         int    `json:"proxyType"`         // Proxy Type：0-http 1-https 2-socks5 3-ssh
}
