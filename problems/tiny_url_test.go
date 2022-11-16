package problems

import (
	"testing"

	. "github.com/hambrosia/algorithms-golang/testhelpers"
)

func TestTinyUrl(t *testing.T) {

	test_url := "https://foosite.com/bar/baz.html"

	shortener := Constructor()
	short_url := shortener.encode(test_url)
	decoded_short_url := shortener.decode(short_url)
	AssertEquals(t, decoded_short_url, test_url)
}
