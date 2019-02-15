package input

import "errors"
import "github.com/reisraff/goinput/input/nodes"
import "github.com/reisraff/goinput/input/interfaces"


type DefaultTypeHandler struct {
    types map[string]nodes.NodeFactory
    _errors []string
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

    return &handler
}

func (self * DefaultTypeHandler) GetType(_type string) (interfaces.NodeInterface, error) {
    if val, ok := self.types[_type]; ok {
        result := val()
        // result.SetTypeAlias(_type)
        result.SetTypeHandler(self)

        return result, nil;
    }

    return nil, errors.New("Type " + _type + " does not exists.")
}

func (self * DefaultTypeHandler) AddError(err string) {
    self._errors = append(self._errors, err)
}

func (self DefaultTypeHandler) GetErrors() []string {
    return self._errors
}
