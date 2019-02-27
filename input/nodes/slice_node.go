package nodes

import "github.com/reisraff/goinput/input/interfaces"
import "reflect"
import "fmt"

func CreateSliceNode() interfaces.NodeInterface {
    node := SliceNode{}
    node.SetRequired(true)

    return &node
}

type SliceNode struct {
    BaseNode
}

func (self *SliceNode) GetValue(field string, value interface{}) interface{} {
    self.CheckConstraints(field, value);

    slice := reflect.ValueOf(value)

    reflection := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(self.GetType())), slice.Len(), slice.Cap())

    for i := 0; i < slice.Len(); i++ {
        obj := self.GetInstantiator().Instantiate(self.GetType(), slice.Index(i).Interface())
        reflection.Index(i).Set(reflect.ValueOf(obj).Elem())
    }

    return reflection.Interface()
}

func (self SliceNode) Walk(input interface{}, parentField string) interface{} {
    if input == nil {
        return input
    }

    result := []interface{}{}

    if (! self.HasChildren()) {
        return input
    }
    
    if "[]interface {}" != reflect.TypeOf(input).String() {
        self.typeHandler.AddError(fmt.Sprintf("value '%v' for field '%s' is invalid", input, parentField))
        return result
    }

    for _, inputItem := range input.([]interface{}) {
        itemResult := make(map[string]interface{})

        for field, node := range self.children {
            if "map[string]interface {}" == reflect.TypeOf(inputItem).String() {
                if _, ok := inputItem.(map[string]interface{})[field]; !ok {
                    if node.IsRequired() {
                        self.typeHandler.AddError(fmt.Sprintf("field '%s' is required", field))
                    }

                    if ! node.HasDefaultValue() {
                        continue
                    }

                    inputItem.(map[string]interface{})[field] = node.GetDefaultValue()
                }

                itemResult[field] = node.GetValue(field, node.Walk(inputItem.(map[string]interface{})[field], field))
            } else {
                self.typeHandler.AddError(fmt.Sprintf("value '%v' for field '%s' is invalid", inputItem, parentField))
                break
            }
        }

        result = append(result, itemResult)
    }

    return result
}