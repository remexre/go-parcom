package parcom_test

import (
	. "github.com/onsi/ginkgo"
	"remexre.xyz/go-parcom"
)

var _ = Describe("AnyOf", func() {
	Describe("abc", func() {
		do(parcom.AnyOf("abc"), anyOfTests)
	})
})

var anyOfTests = []test{
	{"a", "", "a", true},
	{"aa", "", "aa", true},
	{"aaa", "", "aaa", true},
	{"ab", "", "ab", true},
	{"abc", "", "abc", true},
	{"abaa", "", "abaa", true},
	{"abada", "da", "aba", true},
	{"dbada", "dbada", "", false},
}
