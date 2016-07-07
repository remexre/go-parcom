package parcom_test

import (
	"fmt"

	"remexre.xyz/go-parcom"
)

func ExampleAlt() {
	parser := parcom.Alt(
		parcom.Tag("<h1>"),
		parcom.Tag("[header]"),
	)
	remaining, value, ok := parser("[header]")
	fmt.Printf("%#v %#v %#v\n", remaining, value, ok)
	// Output: "" "[header]" true
}
