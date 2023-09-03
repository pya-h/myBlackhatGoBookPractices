package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	City		string	`json:"city"`
	RegionCode	string	`json:"region_code"`
	AreaCode	int		`json:"area_code"`
	Longitude	float32	`json:"longitude"`
	CountryCode3 string	`json:"country_code3"`
	CountryName  string	`json:"country_name"`
	PostalCode	 string	`json:"postal_code"`
	DMACode		 int	`json:"dma_code"`
	CountryCode	 string	`json:"country_code"`
	Latitude	 float32 `json:"latitude"`
}

type Host struct {
	OS			string	`json:"os"`
	Timestamp	string	`json:"timestamp"`
	ISP			string	`json:"isp"`
	ASN			string	`json:"asn"`
	Hostnames	[]string `json:"hostnames"`
	Location	Location `json:"location"`
	IP			int64	 `json:"ip"`
	Domains		[]string `json:"domains"`
	Org			string	 `json:"org"`
	Data		string	 `json:"data"`
	Port		int		 `json:"port"`
	IPString	string	 `ip_str`
}

type HostSearchResponse struct {
	Matches		[]Host	`json:"matches"`
}

func (self *Client) HostSearch(query string) (*HostSearchResponse, error) {
	if res, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BASE_URL, self.apiKey, query)); err == nil {
		defer res.Body.Close()
		var response HostSearchResponse
		if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
			return nil, err
		}
		return &response, nil
	} else {
		return nil, err
	}

}