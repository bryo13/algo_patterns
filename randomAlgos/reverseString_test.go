package randomalgos

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	want := "nai"
	got := Reverse("ian")

	if strings.Compare(want, got) != 0 {
		t.Errorf("wanted %s but got %s", want, got)
	}
}

// go test -fuzz=Fuzz -fuzztime 40s
func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, a string) {
		Reverse(a)
	})
}
