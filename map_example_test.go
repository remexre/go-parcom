package parcom_test

import (
	"fmt"
	"strconv"

	"remexre.xyz/go-parcom"
)

func ExampleMap() {
	parser := parcom.Map(parcom.AnyOf("1234567890"), func(str string) int {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return n
	})
	remaining, value, ok := parser("112358")
	fmt.Printf("%#v %#v %#v\n", remaining, value, ok)
	// Output: "" 112358 true
}
