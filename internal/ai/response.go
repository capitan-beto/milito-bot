package ai

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/genai"
)

///CORREGIR TODOS LOS ERROR HANDLERS DE ACA

func Response(msg string) (string, string, error) {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.WithError(err).Error("error creating client in ln 22!")
		return "", "", err
	}

	model := "gemini-2.0-flash"
	config := ModelConfig

	chat, err := client.Chats.Create(ctx, model, config, nil)
	if err != nil {
		log.WithError(err).Error("error creating chat in ln 31!")
		return "", "", err
	}

	result, err := chat.SendMessage(ctx, genai.Part{Text: msg})
	if err != nil {
		log.WithError(err).Error("error sending message to AI in ln 37!")
		return "", "", err
	}

	var tool map[string]any
	var tname, optType string

	fc := result.FunctionCalls()

	if len(fc) > 0 {
		tool, tname, optType = ToolSwitch(fc[0])
		result, err = chat.SendMessage(ctx, *genai.NewPartFromFunctionResponse(tname, tool))
		if err != nil {
			log.WithError(err).Error("error sending wpp message in ln 50!")
		}

		return result.Text(), optType, nil
	} else {
		//ver si es realmente necesario el no_process
		return result.Text(), "no_process", nil
	}
}
