package node

import "github.com/reisraff/go-input/input/interfaces"

func CreateFloatNode() interfaces.NodeInterface {
    return &FloatNode{}
}

type FloatNode struct {
    BaseNode
}