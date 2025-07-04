package wpp

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

func MarkAsRead(message *models.WebhookMessages, bpnid string) error {

	url := fmt.Sprintf(`https://graph.facebook.com/v22.0/%s/messages`, bpnid)

	var b = models.WhatsappBody{
		MessagingProduct: "whatsapp",
		Status:           "read",
		MessageID:        message.Id,
	}

	jsonBody, err := json.Marshal(b)
	if err != nil {
		log.Error(err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Error(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WPP_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Error(fmt.Errorf("response failed with status: %d and \nbody: %s", res.StatusCode, resBody))
		return err
	}
	if err != nil {
		log.Error(err)
		return err
	}

	fmt.Printf("%s", resBody)
	return nil
}
