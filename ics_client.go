package personio

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type IcsClient struct {
	HttpClient *http.Client
	UserAgent  string
	URL        string
}

func NewClient(httpClient *http.Client, userAgent, icsURL string) *IcsClient {
	return &IcsClient{
		HttpClient: httpClient,
		UserAgent:  userAgent,
		URL:        icsURL,
	}
}

func (c *IcsClient) GetIcsData() ([]byte, error) {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		//log.Print(err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code getting ical: %d", res.StatusCode)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil

}
