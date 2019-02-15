package node

import "github.com/reisraff/go-input/input/interfaces"

func CreateIntNode() interfaces.NodeInterface {
    return &IntNode{}
}

type IntNode struct {
    BaseNode
}