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

type Skeleton interface {
	Skeleton(ss string) string
	Confusable(x, y string) bool
}

func New(confusables map[rune]string) Skeleton {
	if confusables == nil {
		confusables = confusablesMap
	}
	return &skeleton{confusables: confusables}
}

type skeleton struct {
	confusables map[rune]string
}

func (s *skeleton) mapConfusableRunes(ss string) string {
	var buffer bytes.Buffer
	for _, r := range ss {
		replacement, replacementExists := s.confusables[r]
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
func (s *skeleton) Skeleton(ss string) string {
	return norm.NFD.String(
		s.mapConfusableRunes(
			norm.NFD.String(ss)))
}

func (s *skeleton) Confusable(x, y string) bool {
	return s.Skeleton(x) == s.Skeleton(y)
}
