package ai

import "google.golang.org/genai"

var ModelConfig *genai.GenerateContentConfig = &genai.GenerateContentConfig{
	Tools:            []*genai.Tool{Tools},
	Temperature:      genai.Ptr[float32](1.65),
	TopK:             genai.Ptr[float32](2.0),
	TopP:             genai.Ptr[float32](.5),
	MaxOutputTokens:  128,
	ResponseMIMEType: "text/plain",
	SystemInstruction: &genai.Content{
		Parts: []*genai.Part{{Text: `
			Eres una asistenta encargada de responder los usuarios de un emprendimiento de piercing
			
			Tus tareas principales son: 
				-Guardar turnos.
				// -Leer turnos a los clientes'
				// -Eliminar turnos si el cliente lo desea.
				// -Modificar turnos si el cliente asi lo desea

			Responde amablemente y cute :3

			Si el cliente quiere reservar un turno debes preguntarle a que horario

			Porfavor no alucines
		`}},
	},
}
