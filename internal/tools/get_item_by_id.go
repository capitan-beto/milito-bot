package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/capitan-beto/macbot/api"
)

func GetItemByID(id string) map[string]any {

	res, err := http.Get(fmt.Sprintf("http://localhost:8080/client/products/%s", id))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var p *api.GetProductByIdResponse
	json.NewDecoder(res.Body).Decode(&p)

	return map[string]any{
		"Code":  p.Code,
		"Id":    p.Id,
		"Desc":  p.Desc,
		"Price": p.Price,
		"Date":  p.Date,
	}
}
