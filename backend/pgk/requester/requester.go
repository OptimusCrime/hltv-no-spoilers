package requester

import (
	"errors"
	"net/url"

	"github.com/RomainMichau/CycleTLS/cycletls"
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

const (
	BaseUrl = "https://www.hltv.org"
)

type RequestParams struct {
	Url   string
	Query *url.Values
}

func MakeRequest(p *RequestParams) ([]byte, error) {
	requestUrl := BaseUrl + p.Url

	if p.Query != nil {
		requestUrl += "?" + p.Query.Encode()
	}

	client, _ := cloudscraper.Init(false, false)
	options := cycletls.Options{
		Headers: map[string]string{
			"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"accept-encoding":           "gzip, deflate, br, zstd",
			"accept-language":           "nb,en-US;q=0.9,en;q=0.8",
			"cache-control":             "no-cache",
			"referer":                   "https://www.hltv.org",
			"pragma":                    "no-cache",
			"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/112.0.0.0",
			"priority":                  "u=0, i",
			"sec-ch-ua":                 "\"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"138\", \"Opera\";v=\"122\"",
			"sec-ch-ua-mobile":          "?0",
			"sec-ch-ua-platform":        "\"macOS\"",
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "same-origin",
			"sec-fetch-user":            "?1",
			"upgrade-insecure-requests": "1",
		},
		Timeout: 10,
	}

	resp, err := client.Do(requestUrl, options, "GET")

	if err != nil {
		return nil, err
	}

	if resp.Status == 500 {
		return nil, errors.New("server responded with 500")
	}

	// This is a bit silly
	return []byte(resp.Body), nil
}
