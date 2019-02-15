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

    int_types := []string{
        "int8",
        "uint8",
        "byte",
        "int16",
        "uint16",
        "int32",
        "rune",
        "uint32",
        "int64",
        "uint64",
        "int",
        "uint",
        "uintptr",
    }

    float_types := []string{
        "float32",
        "float64",
        "complex64",
        "complex128",
    }

    switch self._type {
        case "int":
            for _, v := range int_types {
                if v == reflectType {
                    return true
                }
            }
        case "float":
            for _, v := range float_types {
                if v == reflectType {
                    return true
                }
            }
        case "numeric":
            numeric_types := append(int_types, float_types...)
            for _, v := range numeric_types {
                if v == reflectType {
                    return true
                }
            }
    }

    return reflectType == self._type
}