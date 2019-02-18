package nodes

import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/constraints"

func CreateBoolNode() interfaces.NodeInterface {
    node := BoolNode{}
    node.SetType("bool")
    node.SetRequired(true)
    node.AddConstraint(constraints.ConstraintType("bool"))

    return &node
}

type BoolNode struct {
    BaseNode
}