package nodes

import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/constraints"

func CreateStringNode() interfaces.NodeInterface {
    node := StringNode{}
    node.SetRequired(true)
    node.AddConstraint(constraints.ConstraintType("string"))

    return &node
}

type StringNode struct {
    BaseNode
}