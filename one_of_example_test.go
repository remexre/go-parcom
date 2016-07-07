package parcom_test

import (
	"fmt"

	"remexre.xyz/go-parcom"
)

func ExampleOneOf() {
	parser := parcom.OneOf("abc")
	remaining, value, ok := parser("abacus")
	fmt.Printf("%#v %#v %#v\n", remaining, value, ok)
	// Output: "bacus" "a" true
}
