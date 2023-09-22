package main

import (
	"artajasa/code"
	"fmt"
)

func main() {
	var words string = "BIODATA INTERVIEW\n\nNAMA: NAMA SAYA\nSEKOLAH : BACHELOR OF DEGREE UNIVERSITY OF\nINDONESIA\nTANGGAL DAN JAM: 7/12/2021 15:22:32\n\nPROGRAMMING LANGUAGES: PHP, GO,\nJAVASCRIPT, SQL, PYTHON, RUBY, CSS, HTML, FRAMEWORK DLL(VUE JS, CODEIGNITER, RUBY ON\nRAILS, JQUERY, BOOTSTRAP, MYSQL, WORDPRESS) GOOGLE SERVICE(FIREBASE, GOOGLE ANALYTICS,\nBIGQUERY) VCS(GIT, GITLAB)\nTERIMA KASIH"

	fmt.Println(code.ConvertWord(words))
}
