//go:generate go run maketables.go > tables.go

package confusables

import (
	"bytes"

	"golang.org/x/text/unicode/norm"
)

// TODO: document casefolding approaches
// (suggest to force casefold strings; explain how to catch paypal - pAypal)
// TODO: DOC you might want to store the Skeleton and check against it later
// TODO: implement xidmodifications.txt restricted characters

func mapConfusableRunes(ss string) string {
	var buffer bytes.Buffer
	for _, r := range ss {
		replacement, replacementExists := confusablesMap[r]
		if replacementExists {
			buffer.WriteString(replacement)
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// Skeleton converts a string to it's "skeleton" form
// as descibed in http://www.unicode.org/reports/tr39/#Confusable_Detection
//   1. Converting X to NFD format
//   2. Successively mapping each source character in X to the target string
//      according to the specified data table
//   3. Reapplying NFD
func Skeleton(s string) string {
	return norm.NFD.String(
		mapConfusableRunes(
			norm.NFD.String(s)))
}

func Confusable(x, y string) bool {
	return Skeleton(x) == Skeleton(y)
}
