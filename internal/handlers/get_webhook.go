package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/capitan-beto/macbot/api"
	log "github.com/sirupsen/logrus"
)

func GetWebhook(w http.ResponseWriter, r *http.Request) {
	req, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Error(err)
		return
	}

	mode := req.Get("hub.mode")
	token := req.Get("hub.verify_token")
	challenge := req.Get("hub.challenge")

	fmt.Println(challenge)
	fmt.Println(mode == "subscribe")
	fmt.Println(token == os.Getenv("WEBHOOK_VERIFY_TOKEN"))
	fmt.Println(os.Getenv("WEBHOOK_VERIFY_TOKEN"))

	if mode == "subscribe" && token == os.Getenv("WEBHOOK_VERIFY_TOKEN") {
		fmt.Println("eaaa")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	} else {
		api.UnauthorizedErrorHandler(w)
		w.Write([]byte("Forbidden"))
	}
}
