package node

import "github.com/reisraff/go-input/input/interfaces"
import "github.com/reisraff/go-input/input/constraints"

func CreateNumericNode() interfaces.NodeInterface {
    node := NumericNode{}
    node.AddConstraint(constraints.ConstraintType("numeric"))

    return &node
}

type NumericNode struct {
    BaseNode
}