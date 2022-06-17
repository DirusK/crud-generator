package models

import (
	"strings"
)

var abbreviations = []string{
	"Url",
	"Id",
	"Bcid",
	"Uri",
}

func replaceAbbreviations(str string) string {
	for _, abbreviation := range abbreviations {
		if strings.Contains(str, abbreviation) {
			return strings.ReplaceAll(str, abbreviation, strings.ToUpper(abbreviation))
		}
	}

	return str
}
