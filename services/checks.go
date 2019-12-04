package services

import (
	"regexp"
	"strings"
)

func isTitle(text string, i int, el *[]string) bool {
	return (text == "Scheda film" && i+1 < len(*el) && (*el)[i+1] != "TOP") ||
		(text == "TOP" && i+1 < len(*el))
}

func isHour(text string) bool {
	match, err := regexp.Match("[0-2][0-9][:][0-5][0-9]", []byte(text))
	if err != nil {
		return false
	}
	return match
}

func isLength(text string) bool {
	return strings.HasPrefix(text, "Durata")
}

func formatLength(text string) string {
	text = strings.ReplaceAll(text, "Durata ", "")
	return strings.ReplaceAll(text, " min.", "")
}
