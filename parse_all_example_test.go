package parcom_test

import (
	"fmt"

	"remexre.xyz/go-parcom"
)

func ExampleParseAll() {
	parser := parcom.Tag("a")
	fmt.Println(parcom.ParseAll("aaaaa", parser))
	// Output: [a a a a a] true
}
