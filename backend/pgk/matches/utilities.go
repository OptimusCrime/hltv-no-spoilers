package matches

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

func parseMatchDate(dateString string) (*time.Time, error) {
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
	reg, _ := regexp.Compile("[^0-9]+")
	return reg.ReplaceAllString(dateString, "")
}

func addLeadingZero(str string) string {
	if len(str) == 1 {
		return "0" + str
	}

	return str
}
