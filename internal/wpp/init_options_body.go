package wpp

import "github.com/capitan-beto/macbot/models"

var InitOptsBody = models.WhatsappBody{
	MessagingProduct: "whatsapp",
	RecipientType:    "individual",
	To:               "EDIT_ME",
	Type:             "interactive",
	Interactive: &models.Interactive{
		Type: "list",
		Header: models.Header{
			Type: "text",
			Text: "¿Que servicio necesitás?",
		},
		Body: models.Body{
			Text: "Elegí entre las opciones rápidas o selecciona otro si estas no te pueden ayudar",
		},
		Footer: models.Footer{
			Text: "Powered by Puzzolana",
		},
		Action: models.Action{
			Button: "Elegir opción",
			Sections: []models.Section{
				{
					Title: "Turnos",
					Rows: []models.Row{
						{
							Id:          "create_turno",
							Title:       "Reservar Turno",
							Description: "Seleccioná la perforación la fecha y hora ideales para vos.",
						},
						{
							Id:          "read_turno",
							Title:       "Ver mi turno",
							Description: "Obtené un recordatorio de turno",
						},
						{
							Id:          "delete_turno",
							Title:       "Cancelar mi turno",
							Description: "Cancelá tu turno abonado. Los reembolsos pueden tardar hasta 24hs.",
						},
						{
							Id:          "update_turno",
							Title:       "Modificar mi turno",
							Description: "Modificá el día u horario de tu turno. Es necesario que ya esté aprobado",
						},
					},
				},
				{
					Title: "Consultas Generales",
					Rows: []models.Row{
						{
							Id:    "read_cares",
							Title: "Cuidados",
						},
						{
							Id:    "write_comment",
							Title: "Dejar una opinión",
						},
						{
							Id:    "read_info",
							Title: "Información previa.",
						},
						{
							Id:          "write_other",
							Title:       "Otra pregunta.",
							Description: "Dejanos un mensaje con tu pregunta y la vamos a responder cuanto antes",
						},
					},
				},
			},
		},
	},
}
