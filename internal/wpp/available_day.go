package wpp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/capitan-beto/macbot/internal/tools"
	"github.com/capitan-beto/macbot/models"
	log "github.com/sirupsen/logrus"
)

func SendAvailableDays(to, bpnid string) error {
	b := AvailableDaysBody

	rows := []models.Row{}

	db, err := tools.CreateConnection()
	if err != nil {
		log.WithError(err).Error("error in ln 11")
		return err
	}

	ad, err := tools.AvailableDaysGetter(db)
	if err != nil {
		log.WithError(err).Error("error getting dates in ln 17!")
		return err
	}

	for _, date := range ad {
		rows = append(rows, models.Row{
			Id:    fmt.Sprintf("pick_date_%s", date),
			Title: date,
		})
	}

	b.Interactive.Action.Sections[0].Rows = rows
	b.To = to

	url := fmt.Sprintf(`https://graph.facebook.com/v22.0/%s/messages`, bpnid)

	reqBody, err := json.Marshal(&b)
	if err != nil {
		log.WithError(err).Error("error marshaling json in ln 40")
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
