package interfaces

type NodeInterface interface {
    Add(string, string, map[string]interface{}) (NodeInterface, error)
    GetValue(string, interface{}) interface{}
    Walk(interface{}, string) interface{}
    HasChildren() bool
    SetTypeHandler(TypeHandlerInterface)
    // AllowNull() bool
    CheckConstraints(field string, value interface{}) []string
    AddConstraint(ConstraintInterface)
    IsRequired() bool
    SetRequired(val bool)
}