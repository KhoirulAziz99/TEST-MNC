package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/KhoirulAziz99/mnc/pkg/utility"
	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {
	host := utility.GetEnv("DB_HOST")
	port := utility.GetEnv("DB_PORT")
	user := utility.GetEnv("DB_USER")
	password := utility.GetEnv("DB_PASSWORD")
	dbname := utility.GetEnv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Print("Succsessfully connected")
	}
	return db, nil
}

func DbClose(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	} else {
		log.Println("Databased closed")
	}
}
