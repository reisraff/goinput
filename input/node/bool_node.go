package node

import "github.com/reisraff/input/input/interfaces"

func CreateBoolNode() interfaces.NodeInterface {
    return &BoolNode{}
}

type BoolNode struct {
    BaseNode
}