package parser

import (
	"github.com/optimuscrime/hltv-no-spoilers/pgk/ttokenizer"
	"golang.org/x/net/html"
	"time"
)

type TeamResult struct {
	Id        int64  `json:"id"`
	Team1     string `json:"team1"`
	Team2     string `json:"team2"`
	EventName string `json:"eventName"`
	MatchType string `json:"type"`
}

type TeamResultGroup struct {
	MatchDate *time.Time   `json:"date"`
	Matches   []TeamResult `json:"matches"`
}

func ParseTeamResults(body string) ([]TeamResultGroup, error) {
	tokenizer := ttokenizer.CreateTokenizerFromString(body)

	var teamResultGroups []TeamResultGroup

	for {
		// Only advance to the next token if we are not starting a new match group, or we will skip the next match
		// group all together
		if !isTeamResultGroupStart(tokenizer) {
			tokenizer.Next()
		}

		tokenType := *tokenizer.TokenType

		switch {
		case tokenType == html.ErrorToken:
			return reverseMatches(teamResultGroups), nil
		default:
			if isTeamResultGroupStart(tokenizer) {
				newMatchGroupResponse, err := parseTeamResultGroup(tokenizer)
				if err != nil {
					return nil, err
				}

				teamResultGroups = append(teamResultGroups, newMatchGroupResponse)
			}
		}
	}
}

func parseTeamResultGroup(tokenizer *ttokenizer.Ttokenizer) (TeamResultGroup, error) {
	teamResultGroup := TeamResultGroup{}

	for {
		tokenizer.Next()
		tokenType := *tokenizer.TokenType
		token := *tokenizer.Token

		switch {
		case tokenType == html.ErrorToken:
			return teamResultGroup, nil
		default:
			attributes := token.Attr
			if len(attributes) == 0 {
				continue
			}

			// Make sure that we stop processing this team result group if we find another team result group
			if isTeamResultGroupStart(tokenizer) {
				return teamResultGroup, nil
			}

			for _, attribute := range attributes {
				if isMatchDate(attribute) {
					matchDate, err := parseTeamResultGroupDate(tokenizer)
					if err != nil {
						return teamResultGroup, err
					}

					teamResultGroup.MatchDate = matchDate
				}

				if isTeamResultStart(attribute) {
					matches, err := parseTeamResults(tokenizer)
					if err != nil {
						return teamResultGroup, nil
					}

					teamResultGroup.Matches = append(teamResultGroup.Matches, *matches)
				}
			}
		}
	}
}

func parseTeamResultGroupDate(tokenizer *ttokenizer.Ttokenizer) (*time.Time, error) {
	tokenizer.Next()
	newTokenType := *tokenizer.TokenType
	newToken := *tokenizer.Token
	if newTokenType == html.ErrorToken {
		return nil, ttokenizer.ErrFailedToParse
	}

	return parseMatchDateStr(html.EscapeString(newToken.String()))
}

func parseTeamResults(tokenizer *ttokenizer.Ttokenizer) (*TeamResult, error) {
	match := &TeamResult{}

	for {
		tokenizer.Next()
		tokenType := *tokenizer.TokenType
		token := *tokenizer.Token
		switch {
		case tokenType == html.ErrorToken:
			return match, nil
		default:
			attributes := token.Attr
			if len(attributes) == 0 {
				continue
			}

			for _, attribute := range attributes {
				if isMatchURL(attribute) {
					matchId, err := parseMatchIdFromUrl(attribute.Val)
					if err != nil {
						return nil, err
					}

					match.Id = matchId
				}

				if isTeamName(attribute, token) {
					teamName, err := tokenizer.GetNextTokenString()
					if err != nil {
						return nil, err
					}

					if match.Team1 == "" {
						match.Team1 = teamName
					} else {
						match.Team2 = teamName
					}
				}

				if isEventName(attribute) {
					eventName, err := tokenizer.GetNextTokenString()
					if err != nil {
						return nil, err
					}

					if len(eventName) == 0 {
						break
					}

					match.EventName = eventName
				}

				if isMatchType(attribute) {
					matchType, err := tokenizer.GetNextTokenString()
					if err != nil {
						return nil, err
					}

					if len(matchType) == 0 {
						break
					}

					match.MatchType = matchType

					return match, nil
				}
			}
		}
	}
}
