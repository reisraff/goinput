package interfaces

type TypeHandlerInterface interface {
    GetType(interface{}, bool) (NodeInterface, error)
    AddError(string)
    GetErrors() []string
    SetDefaultInstantiator(InstantiatorInterface)
    GetDefaultInstantiator() InstantiatorInterface
}