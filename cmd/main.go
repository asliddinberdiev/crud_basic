package main

import (
	"log"

	"github.com/asliddinberdiev/crud_basic/pkg/database"
)

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "password",
	})

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	
}
