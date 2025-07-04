package wpp

import "github.com/capitan-beto/macbot/models"

var AvailableDaysBody = models.WhatsappBody{
	MessagingProduct: "whatsapp",
	RecipientType:    "individual",
	To:               "EDIT_ME",
	Type:             "interactive",
	Interactive: &models.Interactive{
		Type: "list",
		Header: models.Header{
			Type: "text",
			Text: "Elegí el día ideal!",
		},
		Body: models.Body{
			Text: "Seleccioná uno de los días que tenemos disponibles",
		},
		Footer: models.Footer{
			Text: "Powered by Puzzolana",
		},
		Action: models.Action{
			Button: "Elegir opción",
			Sections: []models.Section{
				{
					Title: "Días disponibles",
					Rows:  []models.Row{},
				},
			},
		},
	},
}
