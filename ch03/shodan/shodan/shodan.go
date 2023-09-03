package shodan

const BASE_URL string = "https://api.shodan.io"

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}