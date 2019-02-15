package node

import "github.com/reisraff/go-input/input/interfaces"

func CreateStringNode() interfaces.NodeInterface {
    return &StringNode{}
}

type StringNode struct {
    BaseNode
}