package ai

import (
	"testing"

	"google.golang.org/genai"
)

func TestToolSwitch(t *testing.T) {
	var args = make(map[string]any)
	p := genai.NewContentFromFunctionCall("testFunc", args, "")

	_, name, opType := ToolSwitch(p.Parts[0].FunctionCall)
	if name != "testFunc" {
		t.Fatalf("error! _expected: testFunc, got: %s", name)
	}
	if opType != "no_process" {
		t.Fatalf("error! expected: no_process, got: %s", opType)
	}
	//t.Fatalf("funcName: %s", p.Parts[0].FunctionCall.Name)
}
