package main

type IPInfo struct {
	IP       string       `json:"ip"`
	Location LocationInfo `json:"location"`
	Isp      IspInfo      `json:"isp"`
	Privacy  PrivacyInfo  `json:"privacy"`
	// Abuse    AbuseInfo    `json:"abuse"`
}

type LocationInfo struct {
	City          string  `json:"city"`
	District      string  `json:"district"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"regionCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Postal        string  `json:"postal"`
	Timezone      string  `json:"timezone"`
	Loc           string  `json:"loc"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
}
type ASNInfo struct {
	ASN    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

type IspInfo struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
	Asname string `json:"asname"`
	As     string `json:"as"`
	Isp    string `json:"isp"`
}
type CarrierInfo struct {
	Name string `json:"name"`
	Mcc  string `json:"mcc"`
	Mnc  string `json:"mnc"`
}

type PrivacyInfo struct {
	Mobile  bool   `json:"mobile"`
	VPN     bool   `json:"vpn"`
	Proxy   bool   `json:"proxy"`
	Tor     bool   `json:"tor"`
	Relay   bool   `json:"relay"`
	Hosting bool   `json:"hosting"`
	Service string `json:"service"`
}

type AbuseInfo struct {
	Address string `json:"address"`
	Country string `json:"country"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Network string `json:"network"`
	Phone   string `json:"phone"`
}

type ipInfoResponse struct {
	IP       string      `json:"ip"`
	Hostname string      `json:"hostname"`
	City     string      `json:"city"`
	Region   string      `json:"region"`
	Country  string      `json:"country"`
	Loc      string      `json:"loc"`
	Org      string      `json:"org"`
	Postal   string      `json:"postal"`
	Timezone string      `json:"timezone"`
	ASN      ASNInfo     `json:"asn"`
	Isp      IspInfo     `json:"isp"`
	Carrier  CarrierInfo `json:"carrier"`
	Privacy  PrivacyInfo `json:"privacy"`
	Abuse    AbuseInfo   `json:"abuse"`
}
type ipApiResponse struct {
	IP            string  `json:"query"`
	Hostname      string  `json:"hostname"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	RegionName    string  `json:"regionName"`
	Region        string  `json:"region"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Currency      string  `json:"currency"`
	Timezone      string  `json:"timezone"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Offset        float64 `json:"Offset"`
	Hosting       bool    `json:"hosting"`
	Org           string  `json:"org"`
	Postal        string  `json:"zip"`
	Asname        string  `json:"asname"`
	As            string  `json:"as"`
	Isp           string  `json:"isp"`
}

type IPInfoProvider interface {
	FetchIPInfo(ip string) (IPInfo, error)
}

// ipinfoProvider implements IPInfoProvider for the ipinfo.io service.
type ipinfoProvider struct{}

// ipAPIProvider implements IPInfoProvider for the ip-api.com service.
type ipAPIProvider struct{}
