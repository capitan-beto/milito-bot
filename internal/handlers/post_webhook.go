package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/capitan-beto/macbot/api"
	"github.com/capitan-beto/macbot/internal/ai"
	"github.com/capitan-beto/macbot/internal/wpp"
	"github.com/capitan-beto/macbot/models"
	log "github.com/sirupsen/logrus"
)

func PostWebhook(w http.ResponseWriter, r *http.Request) {
	/// devolver respuesta rapida a esto y luego meter goroutine

	var payload models.PostWebhook
	json.NewDecoder(r.Body).Decode(&payload)

	defer r.Body.Close()

	// fmt.Println(payload.Entry[0].Changes[0].Value.MessagingProduct)

	msgs := payload.Entry[0].Changes[0].Value.Messages
	if len(msgs) <= 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	bpnid := payload.Entry[0].Changes[0].Value.Metadata.PhoneNumberId

	if msgs[0].Type == "text" {

		for _, msg := range msgs {

			aiRes, opType, err := ai.Response(msg.Text.Body)
			if err != nil {
				log.Error(err)
				api.InternalErrorHandler(w)
				return
			}

			if err := wpp.Handler(opType, bpnid, aiRes, &msg); err != nil {
				log.WithError(err).Error("error calling wpp handler in ln 36: ")
			}
		}

	} else if msgs[0].Type == "interactive" {
		aiRes, _, err := ai.Response(msgs[0].Interactive.ListReply.Title)
		if err != nil {
			log.WithError(err).Error("error getting response in ln 53!")
			api.InternalErrorHandler(w)
			return
		}

		if err := wpp.Handler(msgs[0].Interactive.ListReply.Id, bpnid, aiRes, &msgs[0]); err != nil {
			log.WithError(err).Error("error from handler in ln 60!")
		}
		// fmt.Println(msgs[0].From)
		// fmt.Println(msgs[0].Interactive.Type)
		fmt.Println(msgs[0].Interactive.ListReply.Id)
		// fmt.Println(msgs[0].Interactive.ListReply.Title)
	}
}
