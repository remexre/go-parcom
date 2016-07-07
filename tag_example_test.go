package parcom_test

import (
	"fmt"

	"remexre.xyz/go-parcom"
)

func ExampleTag() {
	parser := parcom.Tag("tag")
	fmt.Println(parser("tagtext"))
	// Output: text tag true
}
