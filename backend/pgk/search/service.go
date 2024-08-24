package search

import (
	"encoding/json"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/requester"
	"net/url"
)

type team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type searchResultRawStruct struct {
	Teams []team `json:"teams"`
}

func searchForTeams(name string) (*[]team, error) {
	query := url.Values{}
	query.Add("term", name)

	bodyBytes, err := requester.MakeRequest(&requester.RequestParams{Url: "/search", Query: &query})
	if err != nil {
		return nil, err
	}

	var d []searchResultRawStruct
	if err = json.Unmarshal(bodyBytes, &d); err != nil {
		return nil, err
	}

	if len(d) != 1 {
		return &[]team{}, nil
	}

	return &d[0].Teams, nil
}
