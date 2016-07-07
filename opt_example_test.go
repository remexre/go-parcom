package parcom_test

import (
	"fmt"

	"remexre.xyz/go-parcom"
)

func ExampleOpt() {
	parser := parcom.Opt(parcom.Tag("<>"), "default")
	remaining, value, ok := parser("<><><><><>")
	fmt.Printf("%#v %#v %#v\n", remaining, value, ok)
	// Output: "<><><><>" "<>" true
}
