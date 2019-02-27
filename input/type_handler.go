package input

import "errors"
import "github.com/reisraff/goinput/input/nodes"
import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/instantiators"
import "reflect"
import "fmt"

type DefaultTypeHandler struct {
    types map[string]nodes.NodeFactory
    _errors []string
    defaultInstantiator interfaces.InstantiatorInterface
}

func DefaultTypeHandlerFactory() interfaces.TypeHandlerInterface {
    handler := DefaultTypeHandler{}
    handler.types = map[string]nodes.NodeFactory{
        "mixed": nodes.CreateBaseNode,
        "bool": nodes.CreateBoolNode,
        "string": nodes.CreateStringNode,
        "int": nodes.CreateIntNode,
        "float": nodes.CreateFloatNode,
        "numeric": nodes.CreateNumericNode,
        "datetime": nodes.CreateDateTimeNode,
        // "double": nodes.CreateFloatNode,
        // "object": nodes.CreateObjectNode,
    }

    handler.SetDefaultInstantiator(instantiators.SetInstantiator{})

    return &handler
}

func (self *DefaultTypeHandler) SetDefaultInstantiator(instantiator interfaces.InstantiatorInterface) {
    self.defaultInstantiator = instantiator
}

func (self *DefaultTypeHandler) GetDefaultInstantiator() interfaces.InstantiatorInterface {
    return self.defaultInstantiator
}

func (self *DefaultTypeHandler) GetType(_type interface{}, isCollection bool) (interfaces.NodeInterface, error) {
    typeString := reflect.TypeOf(_type).Kind().String()

    switch typeString {
        case "string" :
            if val, ok := self.types[_type.(string)]; ok {
                result := val()
                result.SetType(_type.(string))
                result.SetTypeHandler(self)

                return result, nil;
            } else {
                result := nodes.CreateDummyNode()
                result.SetTypeHandler(self)

                return result, errors.New("Type " + _type.(string) + " does not exists.")
            }
        case "struct":
            if isCollection {
                result := nodes.CreateSliceNode()
                result.SetType(_type)
                result.SetTypeHandler(self)
                result.SetInstantiator(self.GetDefaultInstantiator())

                return result, nil;
            } else {
                result := nodes.CreateStructNode()
                result.SetType(_type)
                result.SetTypeHandler(self)
                result.SetInstantiator(self.GetDefaultInstantiator())

                return result, nil;
            }
        default:
            result := nodes.CreateDummyNode()
            result.SetTypeHandler(self)

            return result, errors.New("Type " + typeString + " does not exists.")
    }

    return nil, errors.New(fmt.Sprintf("Unexpected error. Inputed type: %T", _type))
}

func (self *DefaultTypeHandler) AddError(err string) {
    self._errors = append(self._errors, err)
}

func (self DefaultTypeHandler) GetErrors() []string {
    return self._errors
}
