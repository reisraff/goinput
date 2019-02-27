package nodes

import "github.com/reisraff/goinput/input/interfaces"

func CreateDummyNode() interfaces.NodeInterface {
    node := DummyNode{}
    node.SetType("dummy")

    return &node
}

type DummyNode struct {
    BaseNode
}