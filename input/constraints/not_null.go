package constraints

import "strings"
import "reflect"

func NewNotNull() * NotNull {
    constraint := NotNull{}
    constraint.SetErrorMessage("Unexpected Empty Content")

    return &constraint
}

type NotNull struct
{
    errorMessage string
}

func (self * NotNull) SetErrorMessage(errorMessage string) {
    self.errorMessage = errorMessage
}

func (self NotNull) GetErrorMessage() string {
    return self.errorMessage
}

func (self *NotNull) Validate(value interface{}) bool {
    if value != nil && reflect.TypeOf(value).String() == "string" {
        value = strings.Trim(value.(string), " ")

        return value.(string) != ""
    }

    return value != nil
}
