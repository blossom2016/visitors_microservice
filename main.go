package main

import (
	"fmt"
	_ "fmt"
	"github.com/dlsniper/go-microservice-webinar/handler"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	dataSource := "postgres://localhost:localhost@%s5432/postgres?sslmode=disable"
	//datasource := dbds
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	db, err := sqlx.Connect("postgres", fmt.Sprintf(dataSource, dbHost))
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", handler.Home(db))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

