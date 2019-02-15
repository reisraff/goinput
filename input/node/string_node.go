package node

import "github.com/reisraff/input/input/interfaces"

func CreateStringNode() interfaces.NodeInterface {
    return &StringNode{}
}

type StringNode struct {
    BaseNode
}