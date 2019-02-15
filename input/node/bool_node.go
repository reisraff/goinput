package node

import "github.com/reisraff/go-input/input/interfaces"

func CreateBoolNode() interfaces.NodeInterface {
    return &BoolNode{}
}

type BoolNode struct {
    BaseNode
}