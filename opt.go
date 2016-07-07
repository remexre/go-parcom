package parcom

// Opt is a parser that makes the parser it wraps optional.
func Opt(parser Parser, defaultValue interface{}) Parser {
	return func(in string) (string, interface{}, bool) {
		remaining, value, ok := parser(in)
		if !ok {
			return in, defaultValue, true
		}
		return remaining, value, ok
	}
}
