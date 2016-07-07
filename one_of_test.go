package parcom_test

import (
	. "github.com/onsi/ginkgo"
	"remexre.xyz/go-parcom"
)

var _ = Describe("OneOf", func() {
	Describe("abc", func() {
		do(parcom.OneOf("abc"), oneOfTests)
	})
})

var oneOfTests = []test{
	{"a", "", "a", true},
	{"aa", "a", "a", true},
	{"aaa", "aa", "a", true},
	{"ab", "b", "a", true},
	{"abc", "bc", "a", true},
	{"abaa", "baa", "a", true},
	{"abada", "bada", "a", true},
	{"dbada", "dbada", "", false},
}
