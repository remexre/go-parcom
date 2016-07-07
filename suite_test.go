package parcom_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"remexre.xyz/go-parcom"

	"testing"
)

func TestGoParcom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoParcom Suite")
}

type test struct {
	data      string
	remaining string
	out       interface{}
	ok        bool
}

func do(parser parcom.Parser, ts []test) {
	for _, t := range ts {
		func(t test) {
			It(fmt.Sprintf(`Parses "%s"`, t.data), func() {
				remaining, out, ok := parser(t.data)

				if t.ok {
					Expect(ok).To(BeTrue(), "ok")
					Expect(out).To(Equal(t.out), "out")
					Expect(remaining).To(Equal(t.remaining), "remaining")
				} else {
					Expect(ok).To(BeFalse(), "ok")
				}
			})
		}(t)
	}
}
