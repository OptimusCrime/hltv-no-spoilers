package matches

import (
	"github.com/optimuscrime/hltv-no-spoilers/pgk/requester"
	"net/url"
)

func findMatchesForTeam(teamId string) (*[]MatchGroup, error) {
	query := url.Values{}
	query.Add("team", teamId)

	bodyBytes, err := requester.MakeRequest(&requester.RequestParams{Url: "/results", Query: &query})
	if err != nil {
		return nil, err
	}

	body := string(bodyBytes)

	return parseResults(body)
}
