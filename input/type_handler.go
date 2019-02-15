package input

import "errors"
import "github.com/reisraff/go-input/input/node"
import "github.com/reisraff/go-input/input/interfaces"


type DefaultTypeHandler struct {
    types map[string]node.NodeFactory
    _errors []string
}

func DefaultTypeHandlerFactory() interfaces.TypeHandlerInterface {
    handler := DefaultTypeHandler{}
    handler.types = map[string]node.NodeFactory{
        "object": node.CreateBaseNode,
        "bool": node.CreateBoolNode,
        "string": node.CreateStringNode,
        "int": node.CreateIntNode,
        "float": node.CreateFloatNode,
        "numeric": node.CreateNumericNode,
        // "double": node.CreateFloatNode,
        // "object": node.CreateObjectNode,
        // "datetime": node.CreateDateTimeNode,
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
