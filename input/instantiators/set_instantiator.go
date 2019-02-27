package instantiators

import "reflect"
import "strings"
import "fmt"

type SetInstantiator struct {
}

func (self SetInstantiator) Instantiate(_type interface{}, value interface{}) interface{} {
    result := reflect.New(reflect.TypeOf(_type))

    for key, value := range value.(map[string]interface{}) {
        method := fmt.Sprintf("Set%s", strings.Title(key))

        reflectMethod := result.MethodByName(method)
        if reflectMethod.IsValid() {
            inputs := make([]reflect.Value, 1)
            inputs[0] = reflect.ValueOf(value)

            reflectMethod.Call(inputs)
        }
    }

    return result.Interface()
}