package nodes

import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/constraints"
import "github.com/reisraff/goinput/input/transformers"

func CreateDateTimeNode() interfaces.NodeInterface {
    node := DateTimeNode{}
    node.SetRequired(true)
    node.AddConstraint(constraints.ConstraintDateTime())
    node.SetTransformer(transformers.DateTimeTransformer{})

    return &node
}

type DateTimeNode struct {
    BaseNode
}