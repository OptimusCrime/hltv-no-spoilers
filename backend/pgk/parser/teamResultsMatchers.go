package parser

import (
	"github.com/optimuscrime/hltv-no-spoilers/pgk/ttokenizer"
	"golang.org/x/net/html"
	"strings"
)

func isTeamResultGroupStart(t *ttokenizer.Ttokenizer) bool {
	if t.Token == nil {
		return false
	}

	attributes := t.Token.Attr
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

func isMatchDate(attr html.Attribute) bool {
	return attr.Key == "class" && strings.Contains(attr.Val, "standard-headline")
}

func isTeamResultStart(attr html.Attribute) bool {
	return attr.Key == "class" && strings.Contains(attr.Val, "result-con")
}

func isMatchURL(attr html.Attribute) bool {
	return attr.Key == "href" && strings.HasPrefix(attr.Val, "/matches/")
}

func isTeamName(attr html.Attribute, t html.Token) bool {
	if attr.Key != "class" || !strings.Contains(attr.Val, "team") {
		return false
	}

	if strings.Contains(attr.Val, "team1") || strings.Contains(attr.Val, "team2") {
		return false
	}

	// Make sure that we only capture divs
	if t.Data != "div" {
		return false
	}

	return true
}

func isEventName(attr html.Attribute) bool {
	return attr.Key == "class" && strings.Contains(attr.Val, "event-name")
}

func isMatchType(attr html.Attribute) bool {
	return attr.Key == "class" && strings.Contains(attr.Val, "map-text")
}
