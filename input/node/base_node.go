package node

import "fmt"
import "reflect"
import "github.com/reisraff/go-input/input/interfaces"

type NodeFactory func() interfaces.NodeInterface

func CreateBaseNode() interfaces.NodeInterface {
    node := BaseNode{}
    node.SetRequired(true)

    return &node
}

type BaseNode struct {
    typeHandler interfaces.TypeHandlerInterface
    children map[string]interfaces.NodeInterface
    constraints []interfaces.ConstraintInterface
    required bool
}

func (self * BaseNode) SetRequired(required bool) {
    self.required = required
}

func (self BaseNode) IsRequired() bool {
    return self.required
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

func (self BaseNode) Walk(input interface{}, parentField string) interface{} {
    if input == nil {
        return input
    }

    result := make(map[string]interface{})

    if (! self.HasChildren()) {
        return input
    }

    for field, node := range self.children {
        if "map[string]interface {}" == reflect.TypeOf(input).String() {
            if _, ok := input.(map[string]interface{})[field]; !ok {
                if node.IsRequired() {
                    self.typeHandler.AddError(fmt.Sprintf("field '%s' is required", field))
                }

                // if ! node.hasDefault() {
                //     continue
                // }

                // input.(map[string]interface{})[field] = node.getDefault()
            }

            fmt.Printf("input[%s] = %v\n", field, input.(map[string]interface{})[field])
            result[field] = node.GetValue(field, node.Walk(input.(map[string]interface{})[field], field))
        } else {
            self.typeHandler.AddError(fmt.Sprintf("value %v for field '%s' is invalid", input, parentField))
            break
        }
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