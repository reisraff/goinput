# GoInput

```go
import "github.com/reisraff/goinput/input"
import "github.com/revel/revel"

type Notification struct {
    Text string
    Url string
}

func (self *Notification) SetText(text string) {
    self.Text = text
}

func (self *Notification) SetUrl(url string) {
    self.Url = url
}

func NotificationCreateDefine(self input.InputResult) {
    object := self.Add("notification", Notification{}, nil)

    object.Add("text", "string", nil)
    object.Add(
        "url",
        "string",
        map[string]interface{}{
            "required": false,
        }
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

    notification := result.GetData("notification").(*Notification)

    data = make(map[string]interface{})
    data["data"] = notification

    return c.RenderJSON(data)
}
```
