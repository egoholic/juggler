package value

type Value struct {
	val  interface{}
	sig  string // helps to understand what is interface{}
	prev *Value
}

func New(v interface{}, sig string, prev *Value) *Value {
	return &Value{v, sig, nil}
}

func (v *Value) Val() interface{} {
	return v.val
}

func (v *Value) Sig() string {
	return v.sig
}

func (v *Value) Prev() *Value {
	return v.prev
}
