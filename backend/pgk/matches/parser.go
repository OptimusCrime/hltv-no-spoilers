package matches

import (
	"errors"
	"golang.org/x/net/html"
	"slices"
	"strings"
)

var (
	ErrorFailedToParse = errors.New("failed to parse match")
)

func parseResults(body string) (*[]MatchGroup, error) {
	tokenizer := html.NewTokenizer(strings.NewReader(body))

	var matchGroups []MatchGroup
	var responseToken *html.Token

	for {
		// Can I please be allowed to call .Token() multiple times...
		if responseToken == nil {
			tokenizer.Next()
		}

		token := getToken(tokenizer, responseToken)

		// There has got to be a better way of doing this
		if responseToken != nil {
			responseToken = nil
		}

		tokenType := token.Type

		switch {
		case tokenType == html.ErrorToken:
			return reverseMatches(&matchGroups), nil
		default:
			if !tokenIsMatchGroup(token) {
				continue
			}

			newMatchGroupResponse, err := parseMatchGroup(tokenizer)
			if err != nil {
				return nil, err
			}

			if newMatchGroupResponse.matchGroup.MatchDate != nil {
				matchGroups = append(matchGroups, *newMatchGroupResponse.matchGroup)
			}

			if newMatchGroupResponse.token != nil {
				responseToken = newMatchGroupResponse.token
			}
		}
	}
}

type parseMatchGroupResponse struct {
	matchGroup *MatchGroup
	token      *html.Token
}

func parseMatchGroup(tokenizer *html.Tokenizer) (*parseMatchGroupResponse, error) {
	matchGroup := &MatchGroup{}

	for {
		tokenType := tokenizer.Next()
		switch {
		case tokenType == html.ErrorToken:
			return &parseMatchGroupResponse{
				matchGroup: matchGroup,
				token:      nil,
			}, nil
		default:
			token := tokenizer.Token()
			attributes := token.Attr
			if len(attributes) == 0 {
				continue
			}

			if tokenIsMatchGroup(token) {
				return &parseMatchGroupResponse{
					matchGroup: matchGroup,
					token:      &token,
				}, nil
			}

			for _, attribute := range attributes {
				// Match date
				if attribute.Key == "class" && strings.Contains(attribute.Val, "standard-headline") {
					tokenType = tokenizer.Next()
					if tokenType == html.ErrorToken {
						return &parseMatchGroupResponse{
							matchGroup: matchGroup,
							token:      nil,
						}, nil
					}

					token = tokenizer.Token()
					matchDate, err := parseMatchDate(html.EscapeString(token.String()))
					if err != nil {
						return nil, err
					}

					matchGroup.MatchDate = matchDate
				}

				// The matches in the group
				if attribute.Key == "class" && strings.Contains(attribute.Val, "result-con") {
					matches, err := parseMatch(tokenizer)
					if err != nil {
						return &parseMatchGroupResponse{
							matchGroup: matchGroup,
							token:      nil,
						}, nil
					}

					matchGroup.Matches = append(matchGroup.Matches, *matches)
				}
			}
		}
	}
}

func parseMatch(tokenizer *html.Tokenizer) (*Match, error) {
	match := &Match{}

	for {
		tokenType := tokenizer.Next()
		switch {
		case tokenType == html.ErrorToken:
			return match, nil
		default:
			token := tokenizer.Token()
			attributes := token.Attr
			if len(attributes) == 0 {
				continue
			}

			for _, attribute := range attributes {
				// Match URL
				if attribute.Key == "href" && strings.HasPrefix(attribute.Val, "/matches/") {
					match.Url = html.UnescapeString(attribute.Val)
				}

				// Teams
				if attribute.Key == "class" && strings.Contains(attribute.Val, "team") {
					// Make sure that we only capture divs
					if token.Data != "div" {
						tokenizer.Next()
						break
					}

					// Bahhh
					if strings.Contains(attribute.Val, "team1") || strings.Contains(attribute.Val, "team2") {
						tokenizer.Next()
						break
					}

					tokenType = tokenizer.Next()
					if tokenType == html.ErrorToken {
						return nil, ErrorFailedToParse
					}

					token = tokenizer.Token()
					teamName := html.EscapeString(token.String())

					if match.Team1 == "" {
						match.Team1 = teamName
					} else {
						match.Team2 = teamName
					}
				}

				// Event name
				if attribute.Key == "class" && strings.Contains(attribute.Val, "event-name") {
					tokenType = tokenizer.Next()
					if tokenType == html.ErrorToken {
						return nil, ErrorFailedToParse
					}

					token = tokenizer.Token()
					eventName := strings.Trim(html.EscapeString(token.String()), "")
					if len(eventName) == 0 {
						break
					}

					match.EventName = eventName
				}

				// Match type
				if attribute.Key == "class" && strings.Contains(attribute.Val, "map-text") {
					tokenType = tokenizer.Next()
					if tokenType == html.ErrorToken {
						return nil, ErrorFailedToParse
					}

					token = tokenizer.Token()
					matchType := strings.Trim(html.EscapeString(token.String()), "")
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

func tokenIsMatchGroup(token html.Token) bool {
	attributes := token.Attr
	if len(attributes) == 0 {
		return false
	}

	for _, attribute := range attributes {
		if attribute.Key == "class" && strings.Contains(attribute.Val, "results-sublist") {
			return true
		}
	}

	return false
}

func getToken(tokenizer *html.Tokenizer, token *html.Token) html.Token {
	if token != nil {
		return *token
	}

	return tokenizer.Token()
}

func reverseMatches(matchGroups *[]MatchGroup) *[]MatchGroup {
	for _, matchGroup := range *matchGroups {
		slices.Reverse(matchGroup.Matches)
	}

	return matchGroups
}
