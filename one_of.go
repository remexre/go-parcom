package parcom

// OneOf is a parser that matches one byte in the slice.
func OneOf(str string) Parser {
	return OneOfFunc(func(target byte) bool {
		for _, b := range []byte(str) {
			if b == target {
				return true
			}
		}
		return false
	})
}

// OneOfFunc is a parser that matches one byte for which the given function
// returns true.
func OneOfFunc(f func(byte) bool) Parser {
	return func(in string) (string, interface{}, bool) {
		if len(in) > 0 && f(in[0]) {
			return in[1:], string(in[0]), true
		}
		return in, nil, false
	}
}
