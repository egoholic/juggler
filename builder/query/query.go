package query

type Query struct {
	name      string
	inputSig  string
	outputSig string
}

func New(name, inputSig, outputSig string) *Query {
	return &Query{name, inputSig, outputSig}
}
