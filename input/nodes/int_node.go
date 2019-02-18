package nodes

import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/constraints"

func CreateIntNode() interfaces.NodeInterface {
    node := IntNode{}
    node.SetRequired(true)
    node.SetType("int")
    node.AddConstraint(constraints.ConstraintType("int"))

    return &node
}

type IntNode struct {
    BaseNode
}