package wpp

import "github.com/capitan-beto/macbot/models"

var AvailableTimesBody = models.WhatsappBody{
	MessagingProduct: "whatsapp",
	RecipientType:    "individual",
	To:               "EDIT_ME",
	Type:             "interactive",
	Interactive: &models.Interactive{
		Type: "list",
		Header: models.Header{
			Type: "text",
			Text: "Seleccioná la hora más cómoda",
		},
		Body: models.Body{
			Text: "Elegí un horario de tu preferencia",
		},
		Footer: models.Footer{
			Text: "Powered by Puzzolana",
		},
		Action: models.Action{
			Button: "Elegír opción",
			Sections: []models.Section{
				{
					Title: "Horarios disponibles",
					Rows:  []models.Row{},
				},
			},
		},
	},
}
