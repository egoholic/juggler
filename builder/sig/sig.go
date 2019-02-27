package sig

type Field struct {
	name string
	typ  *Sig
}

type Sig struct {
	name string
}

func (s *Sig) Name() string {
	return s.name
}

type Namespace struct {
	name   string
	sigs   map[string]*Sig
	fields map[string]*Field
}

func NewNamespace(name string) *Namespace {
	return &Namespace{name, map[string]*Sig{}, map[string]*Field{}}
}

func (ns *Namespace) AddSig(s *Sig) (err error) {
	ns.sigs[s.Name()] = s
	return nil
}
