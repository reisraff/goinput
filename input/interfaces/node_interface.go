package interfaces

type NodeInterface interface {
    Add(string, interface{}, map[string]interface{}) (NodeInterface, error)
    GetValue(string, interface{}) interface{}
    Walk(interface{}, string) interface{}
    HasChildren() bool
    SetTypeHandler(TypeHandlerInterface)
    CheckConstraints(field string, value interface{}) []string
    AddConstraint(ConstraintInterface)
    IsRequired() bool


    SetRequired(bool)
    SetDefaultValue(interface{})
    SetInstantiator(InstantiatorInterface)
    GetInstantiator() InstantiatorInterface
    SetTransformer(TransformerInterface)
    AddConstraints([]ConstraintInterface)
    SetAllowNull(bool)

    HasDefaultValue() bool
    GetDefaultValue() interface{}
    AllowNull() bool

    GetType() interface{}
    SetType(interface{})
}