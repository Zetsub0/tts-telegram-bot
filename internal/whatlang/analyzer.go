package whatlang

import (
	"errors"
	"fmt"
	"github.com/01walid/goarabic"
	"unicode"
)

func Analyze(text string) (string, error) {
	text = goarabic.ToGlyph(text)
	langCount := map[string]int{
		"ru": 0,
		"en": 0,
		"ar": 0,
	}
	for _, char := range text {
		if unicode.IsLetter(char) {
			if unicode.In(char, unicode.Cyrillic) {
				langCount["ru"]++
			} else if unicode.In(char, unicode.Latin) {
				langCount["en"]++
			} else if unicode.In(char, unicode.Arabic) {
				langCount["ar"]++
			}
		}
	}
	var lang string
	var maxVal int
	for key, val := range langCount {
		if val > maxVal {
			lang = key
			maxVal = val
		}
	}
	if maxVal < len([]rune(text))/2 {
		return "", errors.New("can not define language")
	}
	fmt.Println(lang)
	return lang, nil
}
