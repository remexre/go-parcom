package parcom

// ParseAll parses a string as many times as possible, returning a slice of the
// results, along with a success/failure boolean. If the expression is not
// completely parsed, failure will result.
func ParseAll(input string, parser Parser) ([]interface{}, bool) {
	ok := true
	var out []interface{}
	for ok && len(input) > 0 {
		var val interface{}
		input, val, ok = parser(input)
		out = append(out, val)
	}
	return out, len(input) == 0
}
