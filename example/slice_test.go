package slice

import (
	"testing"
)

type SliceTest struct {
	in  string
	out []byte
}

var b = []byte{'1'}
var SliceTests = []SliceTest{
	SliceTest{"TestFinDigits.txt", b},
}

func TestFindDigits(t *testing.T) {
	for _, r := range SliceTests {
		o := FindDigits(r.in)
		if r.out[0] != o[0] {
			t.Errorf("%q expected %q got %q", r.in, r.out, o)
		}
	}
}
