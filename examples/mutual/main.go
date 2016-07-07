package main

import (
	"reflect"

	"remexre.xyz/go-parcom"
)

func A(in string) (string, interface{}, bool) {
	return parcom.Chain(
		parcom.Tag("a"),
		parcom.Opt(B, nil),
	)(in)
}

func B(in string) (string, interface{}, bool) {
	return parcom.Chain(
		parcom.Tag("b"),
		parcom.Opt(A, nil),
	)(in)
}

func assert(b bool) {
	if !b {
		panic("failed assertion")
	}
}

func main() {
	remaining, out, ok := A("abababba")
	assert(ok)
	assert(remaining == "ba")
	assert(reflect.DeepEqual(out, []interface{}{
		"a",
		[]interface{}{
			"b",
			[]interface{}{
				"a",
				[]interface{}{
					"b",
					[]interface{}{
						"a",
						[]interface{}{
							"b",
							nil,
						},
					},
				},
			},
		},
	}))
}
