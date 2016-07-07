package parcom

// Chain is a parser which requires all of a chain of other parsers to match.
func Chain(parsers ...Parser) Parser {
	return func(in string) (string, interface{}, bool) {
		var out []interface{}
		for _, parser := range parsers {
			var val interface{}
			var ok bool
			in, val, ok = parser(in)
			if !ok {
				return in, val, false
			}
			out = append(out, val)
		}
		return in, out, true
	}
}
