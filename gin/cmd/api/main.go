package main

import (
	"fmt"
	"log"
	"note_api/internal/config"
	"note_api/internal/db"
	"note_api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(client)

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
