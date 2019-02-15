package interfaces

type TypeHandlerInterface interface {
    GetType(string) (NodeInterface, error)
    AddError(string)
    GetErrors() []string
}