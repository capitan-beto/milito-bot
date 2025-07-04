package wpp

import (
	"time"

	"github.com/capitan-beto/macbot/internal/tools"
	"github.com/capitan-beto/macbot/models"
	log "github.com/sirupsen/logrus"
)

func Handler(opType, bpnid, aiRes string, msg *models.WebhookMessages) error {
	db, err := tools.CreateConnection()
	if err != nil {
		log.WithError(err).Error("error creating db in wpp handler")
		return err
	}
	/// db methods para manejar historial
	_, err = tools.AddChatToHistory(db, msg.From, msg.Text.Body, time.Now().Format("20060102150405"), "user")

	h, err := tools.GetChatByPhone(db, msg.From)
	if err != nil {
		log.WithError(err).Error("error fetching chat history")
	}

	if opType == "no_process" {
		if err := AnswerMessage(aiRes, msg.From, bpnid); err != nil {
			log.Error(err)
			return err
		}

		if err := SendInitOptionsUser(msg.From, bpnid); err != nil {
			log.WithError(err).Error("error sending init options to user in ln 16: ")
			return err
		}
	} else if opType == "create_turno" {
		if err := AnswerMessage(aiRes, msg.From, bpnid); err != nil {
			log.WithError(err).Error("error answering msg in ln 20!")
			return err
		}

		if err := SendAvailableDays(msg.From, bpnid); err != nil {
			log.WithError(err).Error("error sending free days!")
			return err
		}
	} else if opType[:9] == "pick_date" {
		if err := AnswerMessage(aiRes, msg.From, bpnid); err != nil {
			log.WithError(err).Error("error answering msg in ln 31!")
			return err
		}

		if err := SendAvailableTimes(msg.From, bpnid, msg.Interactive.ListReply.Title); err != nil {
			log.WithError(err).Error("error sendinavailable times in ln 36!")
		}
	} else if opType[:9] == "pick_time" {
		if err := AnswerMessage(aiRes, msg.From, bpnid); err != nil {
			log.WithError(err).Error("error sending confirmation message! ln 40")
			return err
		}

		if err := AnswerMessage("¿Podrías indicarme tu nombre para añadir tu turno?", msg.From, bpnid); err != nil {
			log.WithError(err).Error("error sending confirmation message! ln 40")
			return err
		}
	}

	if err := MarkAsRead(msg, bpnid); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
