package node

import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/constraints"

func CreateNumericNode() interfaces.NodeInterface {
    node := NumericNode{}
    node.SetRequired(true)
    node.AddConstraint(constraints.ConstraintType("numeric"))

    return &node
}

type NumericNode struct {
    BaseNode
}