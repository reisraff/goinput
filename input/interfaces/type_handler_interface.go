package interfaces

type TypeHandlerInterface interface {
    GetType(interface{}) (NodeInterface, error)
    AddError(string)
    GetErrors() []string
    SetDefaultInstantiator(InstantiatorInterface)
    GetDefaultInstantiator() InstantiatorInterface
}