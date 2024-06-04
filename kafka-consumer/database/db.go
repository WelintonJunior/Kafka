package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	user := "root"
	pass := "secret"
	database := "Kafka"
	ip := "localhost"
	port := "5432"

	conn := fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable", user, pass, database, ip, port)
	DB, err = sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}
}
