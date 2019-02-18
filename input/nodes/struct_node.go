package nodes

import "github.com/reisraff/goinput/input/interfaces"
import "reflect"
import "strings"
import "fmt"

func CreateStructNode() interfaces.NodeInterface {
    node := StructNode{}
    node.SetRequired(true)

    return &node
}

type StructNode struct {
    BaseNode
}

func (self *StructNode) GetValue(field string, value interface{}) interface{} {
    self.CheckConstraints(field, value);

    if (self.HasDefaultValue()) {
        object := self.GetDefaultValue();

        for key, value := range value.(map[string]interface{}) {
            method := fmt.Sprintf("Set%s", strings.Title(key))

            inputs := make([]reflect.Value, 1)
            inputs[0] = reflect.ValueOf(value)
            reflect.ValueOf(object).MethodByName(method).Call(inputs)
        }

        return object;
    }

    return self.GetInstantiator().Instantiate(self.GetType(), value);
}