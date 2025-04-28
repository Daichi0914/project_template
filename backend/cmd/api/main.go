package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Daichi0914/project_template/internal/bootstrap"
	"github.com/Daichi0914/project_template/internal/router"
)

func main() {
	dbConn := bootstrap.InitDB()
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatalf("Failed to close DB connection: %v", err)
		}
	}(dbConn)

	r := router.New()
	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
