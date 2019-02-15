# GoInput

```go
import "github.com/reisraff/goinput/input"

func NotificationCreateDefine(self input.InputResult) {
    object := self.Add("field_object", "object", nil)

    object.Add("field_string", "string", nil)
    object.Add(
        "field_numeric",
        "numeric",
        map[string]interface{}{
            "required": false,
        },
    )
    object.Add(
        "field_bool",
        "bool",
        map[string]interface{}{
            "default_value": false,
        },
    )
}

func (c NotificationController) Create() revel.Result {
    handler := input.InputHandler{}
    handler.Configure(input.DefaultTypeHandlerFactory())

    data := make(map[string]interface{})
    c.Params.BindJSON(&data)

    result := handler.Bind(data, NotificationCreateDefine)

    if ! result.IsValid() {
        data = make(map[string]interface{})
        data["errors"] = result.GetErrorsAsString()

        return c.RenderJSON(data)
    }

    data = make(map[string]interface{})
    data["data"] = result.GetData("object")

    return c.RenderJSON(data)
}
```
