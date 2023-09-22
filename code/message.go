package code

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ConvertWord ....
func ConvertWord(words string) (result string) {
	// split with \n
	wordsSplit := strings.Split(words, "\n")
	// array to string
	wordsStr := strings.Join(wordsSplit, " ")
	// split with space
	wordsArr := strings.Split(wordsStr, " ")

	for _, word := range wordsArr {
		switch cases.Lower(language.English).String(word) {
		case "of", "dan":
			result += cases.Lower(language.English).String(word) + " "
		case "php,", "go,", "css,":
			result += strings.ToUpper(word) + " "
		case "interview", "15:22:32":
			result += cases.Title(language.English).String(word) + "\n\n "
		case "sekolah", "tanggal", "terima":
			result += "\n  " + cases.Title(language.English).String(word) + " "
		case "nama:", "jam:", "languages:":
			char := word[:len(word)-1] // remove last character in string
			result += cases.Title(language.English).String(char) + " : "
		default:
			result += cases.Title(language.English).String(word) + " "
		}
	}

	return
}
