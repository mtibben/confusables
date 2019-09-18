package confusables

import (
	"fmt"
	"testing"
)

func TestSkeleton(t *testing.T) {
	confusables := New(nil)

	input := "ρ⍺у𝓅𝒂ן"
	expected := "paypal"
	skeleton := confusables.Skeleton(input)

	if skeleton != expected {
		t.Error(fmt.Sprintf("Skeleton(%s) should result in %s", input, expected))
	}
}

func TestCompareEqual(t *testing.T) {
	confusables := New(nil)

	vectors := [][]string{
		[]string{"ρ⍺у𝓅𝒂ן", "𝔭𝒶ỿ𝕡𝕒ℓ"},
		[]string{"𝖶", "W"},
		[]string{"so̷s", "søs"},
		[]string{"paypal", "paypal"},
		[]string{"scope", "scope"},
		[]string{"ø", "o̷"},
		[]string{"O", "0"},
		[]string{"ν", "v"},
		[]string{"Ι", "l"},
		[]string{"warning", "waming"},
	}

	for _, v := range vectors {
		s1, s2 := v[0], v[1]
		if !confusables.Confusable(s1, s2) {
			t.Errorf("Skeleton strings %+q and %+q were expected to be equal", s1, s2)
		}
	}
}

func TestCompareDifferent(t *testing.T) {
	confusables := New(nil)

	s1 := "Paypal"
	s2 := "paypal"

	if confusables.Confusable(s1, s2) {
		t.Errorf("Skeleton strings %+q and %+q were expected to be different", s1, s2)
	}
}

func BenchmarkSkeletonNoop(b *testing.B) {
	confusables := New(nil)
	s := "skeleton"

	for i := 0; i < b.N; i++ {
		confusables.Skeleton(s)
	}
}

func BenchmarkSkeleton(b *testing.B) {
	confusables := New(nil)
	s := "ѕ𝗄℮|е𝗍ο𝔫"

	for i := 0; i < b.N; i++ {
		confusables.Skeleton(s)
	}
}
