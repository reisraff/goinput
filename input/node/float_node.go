package node

import "github.com/reisraff/go-input/input/interfaces"
import "github.com/reisraff/go-input/input/constraints"

func CreateFloatNode() interfaces.NodeInterface {
    node := FloatNode{}
    node.AddConstraint(constraints.ConstraintType("float"))

    return &node
}

type FloatNode struct {
    BaseNode
}