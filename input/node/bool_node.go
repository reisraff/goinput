package node

import "github.com/reisraff/go-input/input/interfaces"
import "github.com/reisraff/go-input/input/constraints"

func CreateBoolNode() interfaces.NodeInterface {
    node := BoolNode{}
    node.AddConstraint(constraints.ConstraintType("bool"))

    return &node
}

type BoolNode struct {
    BaseNode
}