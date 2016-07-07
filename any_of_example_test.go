package parcom_test

import (
	"fmt"

	"remexre.xyz/go-parcom"
)

func ExampleAnyOf() {
	parser := parcom.AnyOf("abc")
	remaining, value, ok := parser("abacus")
	fmt.Printf("%#v %#v %#v\n", remaining, value, ok)
	// Output: "us" "abac" true
}
