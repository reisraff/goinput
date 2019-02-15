package constraints

import "reflect"

type Type struct {
    _type string
    errorMessage string
}

func ConstraintType(_type string) Type {
    result := Type{}
    result.SetType(_type)
    result.SetErrorMessage("Value does not match type: " + _type)

    return result
}

func (self * Type) SetType(_type string) {
    self._type = _type
}

func (self * Type) SetErrorMessage(errorMessage string) {
    self.errorMessage = errorMessage
}

func (self Type) GetErrorMessage() string {
    return self.errorMessage
}

func (self Type) Validate(value interface{}) bool {
    if value == nil {
        return false
    }

    reflectType := reflect.TypeOf(value).String()

    return reflectType == self._type
}