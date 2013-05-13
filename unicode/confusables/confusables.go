package confusables

import (
	"code.google.com/p/go.text/unicode/norm"
)

// Skeleton converts a string to it's "skeleton" form
// as descibed in http://www.unicode.org/reports/tr39/#Confusable_Detection
func Skeleton(s string) string {

	// 1. Converting X to NFD format
	s = norm.NFD.String(s)

	// 2. Successively mapping each source character in X to the target string
	// according to the specified data table
	var newRunes []rune
	runes := []rune(s)
	for _, char := range runes {
		replacement, exists := confusablesMap[char]
		if exists {
			newRunes = append(newRunes, replacement...)
		} else {
			newRunes = append(newRunes, char)
		}
	}

	// 3. Reapplying NFD
	s = norm.NFD.String(string(newRunes))

	return s
}

func Confusable(x, y string) bool {
	return Skeleton(x) == Skeleton(y)
}
