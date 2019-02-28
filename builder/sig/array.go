package sig

import "fmt"

type Array struct {
	vsig Sig
}

func (a *Array) SigString() string {
	return fmt.Sprintf("[]%s", a.vsig.SigString())
}
