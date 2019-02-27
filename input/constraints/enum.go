package constraints

import "fmt"

func NewEnum(types []interface{}) * Enum {
    constraint := Enum{}
    constraint.SetErrorMessage(fmt.Sprintf("Alowed values '%v'", types))
    constraint.SetTypes(types)

    return &constraint
}

type Enum struct
{
    errorMessage string
    types []interface{}
}

func (self * Enum) SetErrorMessage(errorMessage string) {
    self.errorMessage = errorMessage
}

func (self Enum) GetErrorMessage() string {
    return self.errorMessage
}

func (self * Enum) SetTypes(types []interface{}) {
    self.types = types
}

func (self Enum) GetTypes() []interface{} {
    return self.types
}

func (self *Enum) Validate(value interface{}) bool {
    for _, v := range self.GetTypes() {
        if v == value {
            return true
        }
    }

    self.SetErrorMessage(fmt.Sprintf("Value '%v' not valid. Alowed values '%v'", value, self.GetTypes()))

    return false
}
