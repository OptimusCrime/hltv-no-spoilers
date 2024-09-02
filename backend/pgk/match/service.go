package match

import (
	"github.com/optimuscrime/hltv-no-spoilers/pgk/parser"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/requester"
)

func findMatchVODs(matchId string) ([]parser.VOD, error) {
	bodyBytes, err := requester.MakeRequest(&requester.RequestParams{Url: "/matches/" + matchId + "/foobar", Query: nil})
	if err != nil {
		return nil, err
	}

	body := string(bodyBytes)

	return parser.ParseMatchVODs(body)
}
