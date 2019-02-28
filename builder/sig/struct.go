package sig

import (
	"fmt"
	"strings"
)

type Field struct {
	name string
	sig  Sig
}

func (f *Field) SigString() string {
	return fmt.Sprintf("%s %s", f.name, f.sig.SigString())
}

type Struct struct {
	fields map[string]*Field
}

func (s *Struct) SigString() string {
	acc := strings.Builder{}
	acc.WriteRune('{')
	for _, field := range s.fields {
		acc.WriteString(field.SigString())
		acc.WriteRune('\n')
	}
	acc.WriteRune('}')
	return acc.String()
}
