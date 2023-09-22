package code

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ConvertWord ....
func ConvertWord(words string) (result string) {
	wordsArr := strings.Split(words, " ")
	for _, word := range wordsArr {
		if cases.Lower(language.English).String(word) == "of" || cases.Lower(language.English).String(word) == "dan" {
			result += cases.Lower(language.English).String(word) + " "
		} else {
			result += cases.Title(language.English).String(word) + " "
		}
	}

	return
}
