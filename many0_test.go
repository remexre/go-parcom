package parcom_test

import (
	. "github.com/onsi/ginkgo"
	"remexre.xyz/go-parcom"
)

var _ = Describe("Many0", func() {
	Describe(`Tag("a")`, func() {
		do(parcom.Many0(parcom.Tag("a")), many0Tests)
	})
})

var many0Tests = []test{
	{"a", "", []interface{}{"a"}, true},
	{"aa", "", []interface{}{"a", "a"}, true},
	{"aaa", "", []interface{}{"a", "a", "a"}, true},
	{"ab", "b", []interface{}{"a"}, true},
	{"aaaba", "ba", []interface{}{"a", "a", "a"}, true},
	{"ababa", "baba", []interface{}{"a"}, true},
	{"aabaa", "baa", []interface{}{"a", "a"}, true},
	{"babaa", "babaa", []interface{}(nil), true},
}
