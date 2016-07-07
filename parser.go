package parcom

// Parser is function that parses a string, returning the input that was not
// consumed by parsing, the output of the parser, and a success/failure boolean.
type Parser func(input string) (remainingInput string, out interface{}, ok bool)
