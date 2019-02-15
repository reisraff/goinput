package node

// import "fmt"
import "github.com/reisraff/go-input/input/interfaces"

type NodeFactory func() interfaces.NodeInterface

func CreateBaseNode() interfaces.NodeInterface {
    return &BaseNode{}
}

type BaseNode struct {
    typeHandler interfaces.TypeHandlerInterface
    children map[string]interfaces.NodeInterface
}

func (self * BaseNode) SetTypeHandler(typeHandler interfaces.TypeHandlerInterface) {
    self.typeHandler = typeHandler
}

func (self BaseNode) Add(key string, _type string, options map[string]interface{}) (interfaces.NodeInterface, error) {
    child, err := self.typeHandler.GetType(_type)

    return child, err
}

func (self BaseNode) GetValue(field string, value interface{}) interface{} {
    // if self.AllowNull() && value == nil {
    //     return value
    // }

    // self.checkConstraints(field, value)

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
