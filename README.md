# GoInput

```go
import "github.com/reisraff/goinput/input"

func NotificationCreateDefine(self input.InputResult) {
    object := self.Add("object", "object", nil)

    object.Add("numeric", "numeric", nil)
    object.Add("string", "string", nil)
    object.Add("bool", "bool", nil)
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
