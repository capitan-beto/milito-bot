package tools

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
)

func TestCreateConnection(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Errorf("error loading env")
	}

	if _, err := CreateConnection(); err != nil {
		t.Errorf("error while creating connection in ln 20!")
	}
}

func TestAvailableDaysGetter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	rows := sqlmock.NewRows([]string{"date"}).
		AddRow("2025-03-03").
		AddRow("2025-03-03")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	res, err := AvailableDaysGetter(db)
	if err != nil {
		t.Fatalf("error selecting rows")
	}

	expected := "result: 2025-03-03, 2025-03-03"
	got := fmt.Sprintf("result: %v, %v", res[0], res[1])
	if got != expected {
		t.Fatalf("error! expected %v, got %v", expected, got)
	}
}

func TestAddChatToHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO chat_history").
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = AddChatToHistory(db, "test-data", "1234", "12121212", "user")
	if err != nil {
		t.Fatalf("db procedure returned err %s", err)
	}
}
