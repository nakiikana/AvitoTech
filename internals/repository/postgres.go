package repository

import (
	"fmt"
	"log"
	"tools/internals/cfg"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(config *cfg.Configuration) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.Name, config.DB.Password))
	if err != nil {
		log.Fatalf("Couldn't connect to DB:%v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Printf("Received no response from DB:%v", err)
	}
	return db, nil
}
