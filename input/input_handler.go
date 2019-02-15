package input

import "github.com/reisraff/go-input/input/node"
import "github.com/reisraff/go-input/input/interfaces"

type InputResult struct {
    root interfaces.NodeInterface
    output interface{}
    errors []string
}

func (i * InputResult) Configure(root interfaces.NodeInterface) {
    i.root = root
}

func (i * InputResult) Add(key string, _type string, options map[string]interface{}) interfaces.NodeInterface {
    node, err := i.root.Add(key, _type, options)

    if err != nil {
        i.errors = append(i.errors, err.Error())
    }

    return node
}

func (i * InputResult) GetData(index string) interface{} {
    return i.output.(map[string]interface{})[index]
}

func (i * InputResult) IsValid() bool {
    return len(i.errors) > 0
}

type Define func(InputResult)

type InputHandlerInterface interface {
    Configure(interfaces.TypeHandlerInterface)
    Bind(map[string]interface{})
}

type InputHandler struct {
    typeHandler interfaces.TypeHandlerInterface
}

func (i * InputHandler) Configure(typeHandler interfaces.TypeHandlerInterface) {
    i.typeHandler = typeHandler
}

func (i * InputHandler) Bind(input map[string]interface{}, definer Define) InputResult {
    rootNode := node.CreateBaseNode()
    rootNode.SetTypeHandler(i.typeHandler)

    result := InputResult{}
    result.Configure(rootNode)

    definer(result)

    result.output = result.root.GetValue("root", result.root.Walk(input))

    return result
}

