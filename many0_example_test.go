package parcom_test

import (
	"fmt"

	"remexre.xyz/go-parcom"
)

func ExampleMany0() {
	parser := parcom.Many0(
		parcom.Tag("<>"),
	)
	remaining, value, ok := parser("<><><><><>")
	fmt.Printf("%#v %#v %#v\n", remaining, value, ok)
	// Output: "" []interface {}{"<>", "<>", "<>", "<>", "<>"} true
}
