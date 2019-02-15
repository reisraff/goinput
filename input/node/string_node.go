package node

import "github.com/reisraff/go-input/input/interfaces"
import "github.com/reisraff/go-input/input/constraints"

func CreateStringNode() interfaces.NodeInterface {
    node := StringNode{}
    node.AddConstraint(constraints.ConstraintType("string"))

    return &node
}

type StringNode struct {
    BaseNode
}