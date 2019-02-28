package node

type Value struct {
	val interface{}
	sig string // helps to understand what is interface{}
}

func NewValue(v interface{}, sig string) *Value {
	return &Value{v, sig}
}

type Node struct {
	id    string
	store map[string]*Value
}

func New(id string) *Node {
	return &Node{id, make(map[string]*Value, 100)}
}

func (n *Node) Run() chan *Value {

}

func (n *Node) run(vch chan *Value) {
	select {}
}
