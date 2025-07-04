package wpp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func SendInitOptionsUser(to, bpnid string) error {
	var err error

	url := fmt.Sprintf(`https://graph.facebook.com/v22.0/%s/messages`, bpnid)

	b := InitOptsBody

	b.To = to

	reqBody, err := json.Marshal(b)
	if err != nil {
		log.WithError(err).Error("error marshaling body, ln 19: ")
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.WithError(err).Error("error creating request, ln 26: ")
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WPP_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.WithError(err).Error("error sending http req, line 37: ")
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err = fmt.Errorf("request failed with status: %d and \nbody : %s", res.StatusCode, resBody)
		log.Error(err)
		return err
	}
	if err != nil {
		log.WithError(err).Error("error parsing response, ln 50: ")
	}

	return err
}
