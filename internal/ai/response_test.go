package ai

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestResponse(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("error loading env variables!")
	}

	_, optType, err := Response("this is a test message!")
	if err != nil {
		t.Fatalf("error while getting AI response! %s", err)
	} else if optType != "no_process" {
		t.Fatalf("error getting optType from AI! %s", optType)
	}

}
