//go:generate go run maketables.go > tables.go

package confusables

import (
	"unicode/utf8"

	"code.google.com/p/go.text/unicode/norm"
)

// TODO: document casefolding approaches
// (suggest to force casefold strings; explain how to catch paypal - pAypal)
// TODO: implement tables other than MA
// (is it secure, even if overprocessing, to check only against MA?)
// TODO: DOC you might want to store the Skeleton and check against it later
// TODO: implement xidmodifications.txt restricted characters

// Skeleton converts a string to it's "skeleton" form
// as descibed in http://www.unicode.org/reports/tr39/#Confusable_Detection
func Skeleton(s string) string {

	// 1. Converting X to NFD format
	s = norm.NFD.String(s)

	// 2. Successively mapping each source character in X to the target string
	// according to the specified data table
	for i, w := 0, 0; i < len(s); i += w {
		char, width := utf8.DecodeRuneInString(s[i:])
		replacement, exists := confusablesMap[char]
		if exists {
			s = s[:i] + replacement + s[i+width:]
			w = len(replacement)
		} else {
			w = width
		}
	}

	// 3. Reapplying NFD
	s = norm.NFD.String(s)

	return s
}

func Confusable(x, y string) bool {
	return Skeleton(x) == Skeleton(y)
}
