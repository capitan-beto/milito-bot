package tools

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"google.golang.org/genai"
)

func CreateConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	db.SetMaxOpenConns(5)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Error(pingErr)
		return nil, fmt.Errorf("error checking connection")
	}

	fmt.Println("Connected!")
	return db, err
}

func AvailableDaysGetter(db *sql.DB) ([]string, error) {
	var days []string

	rows, err := db.Query("SELECT * FROM `free_dates` LIMIT 10")
	if err != nil {
		log.WithError(err).Error("error executing query in ln 46!")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d string
		if err := rows.Scan(&d); err != nil {
			log.WithError(err).Error("error scanning rows in ln 55!")
		}
		days = append(days, d)
	}

	if err := rows.Err(); err != nil {
		log.WithError(err).Error("error in ln 61!")
		return nil, err
	}
	db.Close()

	return days, nil
}

func AvailableTimesGetter(db *sql.DB, day string) ([]string, error) {
	var times []string

	rows, err := db.Query("SELECT free_time FROM free_times WHERE date = ?", day)
	if err != nil {
		log.WithError(err).Error("error reading rows in ln 73!")
		return times, err
	}
	defer rows.Close()

	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			log.WithError(err).Error("error reading rows!")
		}
		times = append(times, t)
	}

	if err := rows.Err(); err != nil {
		log.WithError(err).Error("error iterating rows")
		return times, err
	}

	db.Close()
	return times, nil
}

//chat history methods

func AddChatToHistory(db *sql.DB, from, msg, tstamp, role string) (int64, error) {
	res, err := db.Exec("INSERT INTO chat_history(phone_number, msg, tstap, role) VALUES (?, ?, ?)", from, msg, tstamp, role)
	if err != nil {
		log.WithError(err).Error("error inserting chat record")
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.WithError(err).Error("error reading last insert id")
		return 0, err
	}

	return id, nil
}

func GetChatByPhone(db *sql.DB, from string) ([]*genai.Content, error) {
	var history []*genai.Content

	rows, err := db.Query("SELECT msg, role FROM chat_history WHERE phone_number = ?")
	if err != nil {
		log.WithError(err).Error("error reading chat history from db")
		return history, err
	}
	defer rows.Close()

	for rows.Next() {
		var msg, role string

		if err := rows.Scan(&msg, &role); err != nil {
			log.WithError(err).Error("error scanning chat values")
			return history, err
		}

		history = append(history, genai.NewContentFromText(msg, genai.Role(role)))
	}

	if err := rows.Err(); err != nil {
		log.WithError(err).Error("error iterating through rows")
		return history, err
	}

	return history, nil
}
