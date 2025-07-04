package ai

import (
	"github.com/capitan-beto/macbot/internal/tools"
	"google.golang.org/genai"
)

var Tools = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Name:        "getSingleItemById",
			Description: "Get the data of an item using its ID",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"id": {
						Type:        genai.TypeString,
						Description: "Unique ID of the desired item.",
					},
				},
				Required: []string{"id"},
			},
		},
	},
}

func ToolSwitch(p *genai.FunctionCall) (map[string]any, string, string) {
	var apiRes map[string]any
	var optType string
	name := p.Name

	switch name {
	case "getSingleItemById":
		optType = "process"
		if id, ok := p.Args["id"].(string); ok {
			apiRes = tools.GetItemByID(id)
		} else {
			apiRes = tools.GetItemByID(string(id))
		}
	default:
		optType = "no_process"
	}

	return apiRes, name, optType
}
