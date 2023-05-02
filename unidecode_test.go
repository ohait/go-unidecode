package unidecode_test

import (
	"testing"

	"github.com/ohait/go-unidecode"
)

func TestUnidecode(t *testing.T) {
	test := func(in, expected string) {
		out := unidecode.Decode(in)
		if out == expected {
			t.Logf("%q => %q", in, out)
		} else {
			t.Fatalf("%q => expected %q but got %q", in, expected, out)
		}
	}

	test("foo", "foo")
	test("Hello World!", "Hello World!")
	test("úñícøðé", "unicode")
	test("北京", "Bei Jing ")
	test("هر چه", "hr chh")
	test("سلام", "slm")
	test("خانه", "khnh")
}
