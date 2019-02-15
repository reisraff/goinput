package node

import "github.com/reisraff/input/input/interfaces"

func CreateFloatNode() interfaces.NodeInterface {
    return &FloatNode{}
}

type FloatNode struct {
    BaseNode
}