package nodes

import "fmt"
import "reflect"
import "github.com/reisraff/goinput/input/interfaces"

type NodeFactory func() interfaces.NodeInterface

func CreateBaseNode() interfaces.NodeInterface {
    node := BaseNode{}
    node.SetRequired(true)
    node.SetType("object")

    return &node
}

type BaseNode struct {
    typeHandler interfaces.TypeHandlerInterface
    children map[string]interfaces.NodeInterface
    constraints []interfaces.ConstraintInterface
    required bool
    defaultValue interface{}
    allowNull bool
    transformer interfaces.TransformerInterface
    _type interface{}
    instantiator interfaces.InstantiatorInterface
}

func (self *BaseNode) SetInstantiator(instantiator interfaces.InstantiatorInterface) {
    self.instantiator = instantiator
}

func (self BaseNode) GetInstantiator() interfaces.InstantiatorInterface {
    return self.instantiator
}

func (self *BaseNode) SetType(_type interface{}) {
    self._type = _type
}

func (self BaseNode) GetType() interface{} {
    return self._type
}

func (self *BaseNode) SetRequired(required bool) {
    self.required = required
}

func (self BaseNode) IsRequired() bool {
    return self.required
}

func (self *BaseNode) SetTypeHandler(typeHandler interfaces.TypeHandlerInterface) {
    self.typeHandler = typeHandler
}

func (self *BaseNode) Add(key string, _type interface{}, options map[string]interface{}) (interfaces.NodeInterface, error) {
    child, err := self.typeHandler.GetType(_type)

    if options == nil {
        options = make(map[string]interface{})
    }

    if value, ok := options["required"]; ok {
        child.SetRequired(value.(bool))
    }

    if value, ok := options["default_value"]; ok {
        child.SetDefaultValue(value)
    }

    // if value, ok := options["instantiator"]; ok {
    //     child.SetInstantiator(value)
    // }

    if value, ok := options["transformer"]; ok {
        child.SetTransformer(value.(interfaces.TransformerInterface))
    }

    if value, ok := options["constraints"]; ok {
        child.AddConstraints(value.([]interfaces.ConstraintInterface))
    }

    if value, ok := options["allow_null"]; ok {
        child.SetAllowNull(value.(bool))
    }

    if self.children == nil {
        self.children = map[string]interfaces.NodeInterface{}
    }
    self.children[key] = child

    return child, err
}

func (self *BaseNode) GetValue(field string, value interface{}) interface{} {
    if self.AllowNull() && value == nil {
        return value
    }

    for _, err := range self.CheckConstraints(field, value) {
        self.typeHandler.AddError(err)
    }

    if self.transformer != nil {
        return self.transformer.Transform(value)
    }

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

                if ! node.HasDefaultValue() {
                    continue
                }

                input.(map[string]interface{})[field] = node.GetDefaultValue()
            }

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

func (self *BaseNode) SetDefaultValue(defaultValue interface{}) {
    self.required = false // @todo see if its ok

    self.defaultValue = defaultValue
}

func (self *BaseNode) AddConstraints(constraints []interfaces.ConstraintInterface) {
    self.constraints = append(self.constraints, constraints...)
}

func (self *BaseNode) SetAllowNull(allowNull bool) {
    self.allowNull = allowNull
}

func (self BaseNode) HasDefaultValue() bool {
    return self.GetDefaultValue() != nil
}

func (self *BaseNode) GetDefaultValue() interface{} {
    return self.defaultValue
}

func (self *BaseNode) AllowNull() bool {
    return self.allowNull
}

func (self *BaseNode) SetTransformer(transformer interfaces.TransformerInterface) {
    self.transformer = transformer
}