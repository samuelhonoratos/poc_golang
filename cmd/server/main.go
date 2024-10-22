package main

import (
	"log"
	"orcamento/config"
	"orcamento/di/database"
	"orcamento/server"
)

var (
	cfg = config.New()
)

func main() {
	db, err := database.New(cfg)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	server := server.New(db)
	server.Routes()
}
