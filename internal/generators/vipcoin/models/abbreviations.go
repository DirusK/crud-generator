package models

import (
	"strings"
)

// abbreviations store all abbreviations to replace in upper case.
var abbreviations = []string{
	"Url",
	"Id",
	"Bcid",
	"Uri",
}

// replaceAbbreviations function to check if string contains abbreviation letters and replace them with upper case.
func replaceAbbreviations(str string) string {
	for _, abbreviation := range abbreviations {
		if strings.Contains(str, abbreviation) {
			return strings.ReplaceAll(str, abbreviation, strings.ToUpper(abbreviation))
		}
	}

	return str
}
