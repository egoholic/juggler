package sig

import "fmt"

type Map struct {
	ksig Sig
	vsig Sig
}

func (m *Map) SigString() string {
	return fmt.Sprintf("map[%s]%s", m.ksig.SigString(), m.vsig.SigString())
}
