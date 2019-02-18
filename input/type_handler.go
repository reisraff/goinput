package input

import "errors"
import "github.com/reisraff/goinput/input/nodes"
import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/instantiators"
import "reflect"

type DefaultTypeHandler struct {
    types map[string]nodes.NodeFactory
    _errors []string
    defaultInstantiator interfaces.InstantiatorInterface
}

func DefaultTypeHandlerFactory() interfaces.TypeHandlerInterface {
    handler := DefaultTypeHandler{}
    handler.types = map[string]nodes.NodeFactory{
        "object": nodes.CreateBaseNode,
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

func (self *DefaultTypeHandler) GetType(_type interface{}) (interfaces.NodeInterface, error) {
    typeString := reflect.TypeOf(_type).Kind().String()

    switch typeString {
        case "string" :
            if val, ok := self.types[_type.(string)]; ok {
                result := val()
                result.SetType(_type.(string));
                result.SetTypeHandler(self)

                return result, nil;
            }
        case "struct":
            result := nodes.CreateStructNode();
            result.SetType(_type);
            result.SetTypeHandler(self);
            result.SetInstantiator(self.GetDefaultInstantiator());

            return result, nil;
    }

    return nil, errors.New("Type " + typeString + " does not exists.")
}

func (self *DefaultTypeHandler) AddError(err string) {
    self._errors = append(self._errors, err)
}

func (self DefaultTypeHandler) GetErrors() []string {
    return self._errors
}
