package interfaces

type ConstraintInterface interface {
    GetErrorMessage() string
    Validate(interface{}) bool
}