package parcom_test

import (
	"fmt"

	"remexre.xyz/go-parcom"
)

func ExampleChain() {
	parser := parcom.Map(parcom.Chain(
		parcom.Tag("<h"),
		parcom.AnyOf("123456"),
		parcom.Tag(">"),
	), func(i []string) string {
		var out string
		for _, str := range i {
			out += str
		}
		return out
	})
	remaining, value, ok := parser("<h1>")
	fmt.Printf("%#v %#v %#v\n", remaining, value, ok)
	// Output: "" "<h1>" true
}
