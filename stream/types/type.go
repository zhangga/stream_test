package types

type (
	T interface{}
	R interface{}

	// Predicate Represents a predicate (boolean-valued function) of one argument.
	Predicate func(e T) bool

	// Consumer
	Consumer func(e T)
)
