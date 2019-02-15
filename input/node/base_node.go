package node

import "github.com/reisraff/go-input/input/interfaces"

type NodeFactory func() interfaces.NodeInterface

func CreateBaseNode() interfaces.NodeInterface {
    return &BaseNode{}
}

type BaseNode struct {
    typeHandler interfaces.TypeHandlerInterface
    children map[string]interfaces.NodeInterface
    constraints []interfaces.ConstraintInterface
}

func (self * BaseNode) SetTypeHandler(typeHandler interfaces.TypeHandlerInterface) {
    self.typeHandler = typeHandler
}

func (self * BaseNode) Add(key string, _type string, options map[string]interface{}) (interfaces.NodeInterface, error) {
    child, err := self.typeHandler.GetType(_type)

    if self.children == nil {
        self.children = map[string]interfaces.NodeInterface{}
    }
    self.children[key] = child

    return child, err
}

func (self * BaseNode) GetValue(field string, value interface{}) interface{} {
    // if self.AllowNull() && value == nil {
    //     return value
    // }

    _errors := self.CheckConstraints(field, value)
    for _, err := range _errors {
        self.typeHandler.AddError(err)
    }

    // if (self.transformer) {
    //     return self.transformer.transform(value)
    // }

    return value
}

func (self BaseNode) Walk(input interface{}) interface{} {
    result := make(map[string]interface{})

    if (! self.HasChildren()) {
        return input
    }

    for field, node := range self.children {
        // if value, ok := input.(map[string]interface{})[field]; ok {
        //     if (node.isRequired()) {
        //         throw new RequiredFieldException(field)
        //     }

        //     if ! node.hasDefault() {
        //         continue
        //     }

        //     value = node.getDefault()
        // }

        result[field] = node.GetValue(field, node.Walk(input.(map[string]interface{})[field]))
    }

    return result
}

func (self BaseNode) HasChildren() bool {
    return len(self.children) > 0
}

func (self *BaseNode) AddConstraint(constraint interfaces.ConstraintInterface) {
    self.constraints = append(self.constraints, constraint)
}

func (self BaseNode) CheckConstraints(field string, value interface{}) []string {
    var _errors []string

    for _, constraint := range self.constraints {
        if ! constraint.Validate(value) {
            _errors = append(_errors, constraint.GetErrorMessage())
        }
    }

    return _errors
}