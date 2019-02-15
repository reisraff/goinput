package constraints

import "reflect"
import "time"

type DateTime struct {
    errorMessage string
}

func ConstraintDateTime() DateTime {
    result := DateTime{}
    result.SetErrorMessage("Invalid date/time format")

    return result
}

func (self * DateTime) SetErrorMessage(errorMessage string) {
    self.errorMessage = errorMessage
}

func (self DateTime) GetErrorMessage() string {
    return self.errorMessage
}

func (self DateTime) Validate(value interface{}) bool {
    if "string" != reflect.TypeOf(value).String() {
        return false
    }

    _, err := time.Parse("2006-01-02 15:04:05", value.(string))

    return err == nil
}