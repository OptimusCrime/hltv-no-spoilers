package requester

import (
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
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

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if p.Query != nil {
		req.URL.RawQuery = p.Query.Encode()
	}

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "nb,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("referer", "https://www.hltv.org")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/112.0.0.0")
	req.Header.Add("priority", "u=0, i")

	// Shamelessly stolen from: https://github.com/sweetbbak/go-cloudflare-bypass/blob/main/reqwest/reqwest.go
	// hltv.org is protected behind Cloudflare, and the default Go HTTP client was missing a bunch
	// of cipher suites... Enabling them fixed the issue.
	tlsConfig := http.DefaultTransport.(*http.Transport).TLSClientConfig

	client := &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: 30 * time.Second,
			DisableKeepAlives:   false,

			TLSClientConfig: &tls.Config{
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_AES_128_GCM_SHA256,
					tls.VersionTLS13,
					tls.VersionTLS10,
				},
			},
			DialTLS: func(network, addr string) (net.Conn, error) {
				return tls.Dial(network, addr, tlsConfig)
			},
		}}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = resp.Body.Close(); err != nil {
		return nil, err
	}

	if resp.StatusCode == 500 {
		return nil, errors.New("server responded with 500")
	}

	return bodyBytes, nil
}
