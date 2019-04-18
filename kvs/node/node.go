package node

import "github.com/egoholic/juggler/kvs/value"

type Node struct {
	id    string
	store map[string]*value.Value
}

func New(id string) *Node {
	return &Node{id, make(map[string]*value.Value, 100)}
}

type SetValResult struct {
}

func (n *Node) SetVal(key string, val interface{}, sig string) {
	prevValue := n.store[key]
	newValue := value.New(val, sig, prevValue)
	n.store[key] = newValue
}

func (n *Node) GetVal(key string) (interface{}, string) {
	val := n.store[key]
	return val.Val(), val.Sig()
}
