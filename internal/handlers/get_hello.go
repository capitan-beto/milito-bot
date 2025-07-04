package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/capitan-beto/macbot/models"
	log "github.com/sirupsen/logrus"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	url := "https://graph.facebook.com/v22.0/592213860647160/messages"

	var body = models.WhatsappBody{
		MessagingProduct: "whatsapp",
		To:               "543564663285",
		Type:             "template",
		Template: &models.Template{
			Name: "hello_world",
			Language: models.Language{
				Code: "en_US",
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Error(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Error(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WPP_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Error(fmt.Errorf("response failed with status: %d and \nbody: %s", res.StatusCode, resBody))
		return
	}
	if err != nil {
		log.Error(err)
	}

	fmt.Printf("%s", resBody)
}
