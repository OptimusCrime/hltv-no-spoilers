package parser

import (
	"errors"
	"golang.org/x/net/html"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var matchIdExp = regexp.MustCompile(`(?m)^/matches/([0-9]+).*$`)
var dateExp = regexp.MustCompile("[^0-9]+")

func parseMatchIdFromUrl(url string) (int64, error) {
	escapedUrl := html.UnescapeString(url)
	matches := matchIdExp.FindStringSubmatch(escapedUrl)

	if len(matches) != 2 {
		return 0, errors.New("could not find match id: " + escapedUrl)
	}

	return strconv.ParseInt(matches[1], 10, 0)
}

// lmao
func parseMatchDateStr(dateString string) (*time.Time, error) {
	cleanString := strings.Split(strings.Trim(strings.Replace(dateString, "Results for", "", 1), " "), " ")

	if len(cleanString) != 3 {
		return nil, errors.New("failed to parse match MatchDate")
	}

	month, err := parseMonth(cleanString[0])
	if err != nil {
		return nil, err
	}

	date := parseDate(cleanString[1])

	year := cleanString[2]
	matchDate, err := time.Parse(time.DateOnly, year+"-"+addLeadingZero(month)+"-"+addLeadingZero(date))
	if err != nil {
		return nil, err
	}

	return &matchDate, nil
}

// this is 10/10 genius code
func parseMonth(month string) (string, error) {
	switch strings.ToLower(month) {
	case "january":
		return "1", nil
	case "february":
		return "2", nil
	case "march":
		return "3", nil
	case "april":
		return "4", nil
	case "may":
		return "5", nil
	case "june":
		return "6", nil
	case "july":
		return "7", nil
	case "august":
		return "8", nil
	case "september":
		return "9", nil
	case "october":
		return "10", nil
	case "november":
		return "11", nil
	case "december":
		return "12", nil
	}

	return "", errors.New("failed to parse month")
}

func parseDate(dateString string) string {
	return dateExp.ReplaceAllString(dateString, "")
}

func addLeadingZero(str string) string {
	if len(str) == 1 {
		return "0" + str
	}

	return str
}

func reverseMatches(matchGroups []TeamResultGroup) []TeamResultGroup {
	for _, matchGroup := range matchGroups {
		slices.Reverse(matchGroup.Matches)
	}

	return matchGroups
}
