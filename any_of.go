package parcom

// AnyOf is a parser that matches bytes in the slice, as many as possible.
func AnyOf(str string) Parser {
	bytes := []byte(str)
	return AnyOfFunc(func(target byte) bool {
		for _, b := range bytes {
			if target == b {
				return true
			}
		}
		return false
	})
}

// AnyOfFunc is a parser that matches bytes for which the given function returns
// true, as many as possible.
func AnyOfFunc(f func(byte) bool) Parser {
	return func(in string) (string, interface{}, bool) {
		i := 0
		for i < len(in) {
			if !f(in[i]) {
				break
			}
			i++
		}
		if i == 0 {
			return in, nil, false
		}
		return in[i:], in[:i], true
	}
}
