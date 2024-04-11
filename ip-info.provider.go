package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (p ipinfoProvider) FetchIPInfo(ip string) (IPInfo, error) {
	apiToken := os.Getenv("API_TOKEN")

	resp, err := http.Get("https://ipinfo.io/" + ip + "/json" + "?token=" + apiToken)
	fmt.Printf("Requested URL: %s\n", ip)
	return parseIPInfoResponse(resp, err)
}

// parseIPInfoResponse handles the common logic for parsing the response from an IP information service.
func parseIPInfoResponse(resp *http.Response, err error) (IPInfo, error) {
	var apiResult ipInfoResponse
	var result IPInfo
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(body, &apiResult); err != nil {
		return result, err
	}

	// Convert ipApiResponse to IPInfo
	result = convertApiResponseToIpInfo(apiResult)

	return result, nil
}

func convertApiResponseToIpInfo(apiResponse ipInfoResponse) IPInfo {
	return IPInfo{
		IP: apiResponse.IP,
		Isp: IspInfo{
			Name:   apiResponse.ASN.ASN,
			Domain: apiResponse.ASN.Domain,
			Asname: apiResponse.ASN.Name,
			Type:   apiResponse.ASN.Type,
		},
		Privacy: PrivacyInfo{
			Mobile:  apiResponse.Privacy.Mobile,
			VPN:     apiResponse.Privacy.VPN,
			Proxy:   apiResponse.Privacy.Proxy,
			Tor:     apiResponse.Privacy.Tor,
			Relay:   apiResponse.Privacy.Relay,
			Hosting: apiResponse.Privacy.Hosting,
		},
		Location: LocationInfo{
			City:     apiResponse.City,
			Region:   apiResponse.Region,
			Country:  apiResponse.Country,
			Postal:   apiResponse.Postal,
			Timezone: apiResponse.Timezone,
			Loc:      apiResponse.Loc,
		},
	}
}
