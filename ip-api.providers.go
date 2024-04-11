package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (p ipAPIProvider) FetchIPInfo(ip string) (IPInfo, error) {
	resp, err := http.Get("http://ip-api.com/json/" + ip + "?fields=66846719")
	fmt.Printf("Requested URL: %s\n", ip)
	return parseIPAPIResponse(resp, err)

}
func parseIPAPIResponse(resp *http.Response, err error) (IPInfo, error) {
	var apiResult ipApiResponse
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
	result = convertApiResponseToIpApi(apiResult)

	return result, nil
}

func convertApiResponseToIpApi(apiResponse ipApiResponse) IPInfo {
	return IPInfo{
		IP: apiResponse.IP,
		Isp: IspInfo{
			Name:   apiResponse.Org,
			Asname: apiResponse.Asname,
			As:     apiResponse.As,
			Isp:    apiResponse.Isp,
		},
		Privacy: PrivacyInfo{
			Mobile:  apiResponse.Mobile,
			VPN:     false,
			Proxy:   apiResponse.Proxy,
			Tor:     false,
			Relay:   false,
			Hosting: apiResponse.Hosting,
		},
		Location: LocationInfo{
			City:          apiResponse.City,
			District:      apiResponse.District,
			Region:        apiResponse.RegionName,
			RegionCode:    apiResponse.Region,
			Country:       apiResponse.Country,
			CountryCode:   apiResponse.CountryCode,
			Continent:     apiResponse.Continent,
			ContinentCode: apiResponse.ContinentCode,
			Postal:        apiResponse.Postal,
			Timezone:      apiResponse.Timezone,
			Loc:           fmt.Sprintf("%f,%f", apiResponse.Lat, apiResponse.Lon),
			Lat:           apiResponse.Lat,
			Lon:           apiResponse.Lon,
		},
	}
}
