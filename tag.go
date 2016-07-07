package parcom

// Tag is a parser that matches a literal exactly.
func Tag(tag string) Parser {
	return func(in string) (string, interface{}, bool) {
		if len(tag) <= len(in) && string(tag) == in[:len(tag)] {
			return in[len(tag):], string(tag), true
		}
		return in, nil, false
	}
}
