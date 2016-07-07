package parcom

// Alt is a parser which allows any of the contained parsers to match, giving
// precedence to the first.
func Alt(parsers ...Parser) Parser {
	return func(in string) (string, interface{}, bool) {
		for _, parser := range parsers {
			remaining, val, ok := parser(in)
			if ok {
				return remaining, val, ok
			}
		}
		return in, nil, false
	}
}
