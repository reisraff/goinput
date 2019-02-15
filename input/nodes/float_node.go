package nodes

import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/constraints"

func CreateFloatNode() interfaces.NodeInterface {
    node := FloatNode{}
    node.SetRequired(true)
    node.AddConstraint(constraints.ConstraintType("float"))

    return &node
}

type FloatNode struct {
    BaseNode
}