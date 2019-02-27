package command

type Command struct {
	name      string
	inputSig  string
	outputSig string
}

func New(name, inputSig, outputSig string) *Command {
	return &Command{name, inputSig, outputSig}
}
