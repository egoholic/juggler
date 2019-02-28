package sig

import "fmt"

type Sig interface {
	SigString() string
}

type Namespace struct {
	name string
	sigs map[string]map[string]Sig
}

func NewNamespace(name string) *Namespace {
	return &Namespace{name, map[string]map[string]Sig{}}
}

func (ns *Namespace) AddSig(sig Sig, alias string, version string) (err error) {
	key := alias
	if key == "" {
		key = sig.SigString()
	}

	v, ok := ns.sigs[key]
	if !ok {
		m := map[string]Sig{}
		m[version] = sig
		ns.sigs[key] = m
	} else {
		v[version] = sig
	}

	return nil
}

func (ns *Namespace) GetSigs(alias string) (sigs map[string]Sig, err error) {
	sigs, ok := ns.sigs[alias]
	if !ok {
		return sigs, nil
	}

	return nil, fmt.Errorf("Alias: `%s` does not exist!", alias)
}

func (ns *Namespace) GetSig(alias, version string) (sig Sig, err error) {
	sigs, err := ns.GetSigs(alias)
	if err != nil {
		return nil, err
	}

	sig, ok := sigs[version]
	if !ok {
		return nil, fmt.Errorf("Version: `%s` does not exist!", version)
	}

	return sig, nil
}
