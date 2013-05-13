package confusables

import (
	"fmt"
	"testing"
)

func TestSkeleton(t *testing.T) {
	s := "ρ⍺у𝓅𝒂ן"
	expected := "paypal"
	skeleton := Skeleton(s)

	if skeleton != expected {
		t.Error(fmt.Sprintf("Skeleton(%s) should result in %s", s, expected))
	}
}

func TestCompare(t *testing.T) {
	s1 := "ρ⍺у𝓅𝒂ן"
	s2 := "𝔭𝒶ỿ𝕡𝕒ℓ"

	if !Confusable(s1, s2) {
		t.Error("Skeleton strings were expected to be equal")
	}
}
