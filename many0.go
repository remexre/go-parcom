package parcom

// Many0 is a parser which repeatedly uses its child parser, zero or more times.
func Many0(child Parser) Parser {
	return func(in string) (string, interface{}, bool) {
		var out []interface{}
		for {
			nextIn, val, ok := child(in)
			if !ok {
				break
			}
			in = nextIn
			out = append(out, val)
		}
		return in, out, true
	}
}
