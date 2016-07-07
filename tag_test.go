package parcom_test

import (
	. "github.com/onsi/ginkgo"

	"remexre.xyz/go-parcom"
)

var _ = Describe("Tag", func() {
	do(parcom.Tag("tag"), tagTests)
})

var tagTests = []test{
	{"tag", "", "tag", true},
	{"tagtext", "text", "tag", true},
	{"texttag", "texttag", nil, false},
}
