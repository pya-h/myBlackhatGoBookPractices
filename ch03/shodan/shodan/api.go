package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	QueryCredits int `json:"query_credits"`
	ScanCredits int `json:"scan_credits"`
	Telnet bool `json:"telnet"`
	Plan string `json:"plan"`
	HTTPS bool `json:"https"`
	Unlocked bool `json:"unlocked"`
}

func (self *Client) GetApiInfo() (*APIInfo, error) {
	if res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BASE_URL, self.apiKey)); err == nil {
		defer res.Body.Close()
		var response APIInfo
		if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
			return nil, err
		}
		return &response, nil
	} else {
		return nil, err
	}
}
