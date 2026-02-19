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
	client, _ , err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("DB error")
	}
	defer func (){
		if err := db.Disconnect(client); err != nil{
			log.Printf("mongo disconnected error:%v",err)
		}
	}()

	router := server.NewRouter()

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed")
	}

}