package profile

type Profile struct {
	Name              string            `json:"Name,omitempty"`
	GlobalProxy       string            `json:"GlobalProxy"`
	RouteOrder        string            `json:"RouteOrder,omitempty"`
	RemoteDNSType     string            `json:"RemoteDNSType"`
	RemoteDNSDomain   string            `json:"RemoteDNSDomain"`
	RemoteDNSIP       string            `json:"RemoteDNSIP"`
	DomesticDNSType   string            `json:"DomesticDNSType"`
	DomesticDNSDomain string            `json:"DomesticDNSDomain"`
	DomesticDNSIP     string            `json:"DomesticDNSIP"`
	Geoipurl          string            `json:"Geoipurl"`
	Geositeurl        string            `json:"Geositeurl"`
	LastUpdated       string            `json:"LastUpdated,omitempty"`
	DnsHosts          map[string]string `json:"DnsHosts"`
	DirectSites       []string          `json:"DirectSites"`
	DirectIp          []string          `json:"DirectIp"`
	ProxySites        []string          `json:"ProxySites,omitempty"`
	ProxyIp           []string          `json:"ProxyIp,omitempty"`
	BlockSites        []string          `json:"BlockSites,omitempty"`
	BlockIp           []string          `json:"BlockIp,omitempty"`
	DomainStrategy    string            `json:"DomainStrategy"`
	FakeDNS           string            `json:"FakeDNS"`
	UseChunkFiles     string            `json:"UseChunkFiles"`
}

func baseProfile(name string) *Profile {
	return &Profile{
		Name:              name,
		GlobalProxy:       "true",
		RouteOrder:        "block-proxy-direct",
		RemoteDNSType:     "DoH",
		RemoteDNSDomain:   "https://dns.google/dns-query",
		RemoteDNSIP:       "8.8.8.8",
		DomesticDNSType:   "DoU",
		DomesticDNSDomain: "dns.yandex",
		DomesticDNSIP:     "77.88.8.8",
		Geoipurl:          "https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geoip.dat",
		Geositeurl:        "https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geosite.dat",
		FakeDNS:           "false",
		UseChunkFiles:     "true",
		DnsHosts: map[string]string{
			"dns.google": "8.8.8.8",
			"dns.yandex": "77.88.8.8",
		},
		DomainStrategy: "IPIfNonMatch",
	}
}

func NewProfile(name string, userDirectSites []string, userBlockSites []string, userDirectIPs []string) *Profile {
	p := baseProfile(name)
	p.DirectSites = append([]string{
		"geosite:ru",
		"geosite:geolocation-ru",
		"geosite:category-gov-ru",
		"domain:vk.com",
		"domain:vk.me",
		"domain:userapi.com",
		"domain:vkuseraudio.net",
		"domain:vkuservideo.net",
		"domain:vkvideo.ru",
		"domain:mail.ru",
		"domain:ok.ru",
		"domain:max.ru",
		"domain:yandex.ru",
		"domain:yandex.net",
		"domain:yandex.com",
		"domain:ya.ru",
		"domain:yastatic.net",
		"domain:yandexcloud.net",
		"domain:cdnvideo.ru",
	}, userDirectSites...)
	p.BlockSites = append([]string{"geosite:category-ads-all"}, userBlockSites...)
	p.DirectIp = append([]string{
		"geoip:ru",
		"geoip:private",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}, userDirectIPs...)
	return p
}

func NewStrictProfile(name string) *Profile {
	p := baseProfile(name)
	blockSites := []string{
		"geosite:category-ads-all",
	}
	directIPs := []string{
		"geoip:private",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}
	p.DirectSites = nil
	p.BlockSites = blockSites
	p.DirectIp = directIPs
	return p
}

func NewBypassAllProfile(name string) *Profile {
	p := baseProfile(name)
	p.DirectSites = nil
	p.BlockSites = nil
	p.DirectIp = nil
	return p
}
