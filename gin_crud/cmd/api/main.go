package main

import (
	"fmt"
	config "gin_crud/internal"
	"gin_crud/internal/db"
	"gin_crud/internal/server"
	"log"
)

func main() {
	cfg, error := config.Load()
	if error != nil {
		log.Fatal("Config error", error)
	}

	connection, database, error := db.Connect(cfg)
	if error != nil {
		log.Fatal("Db error", error)
	}

	defer func () {
		if error := db.DisConnect(connection); error != nil {
			log.Printf("Mongo disconnection error: %v", error)
		}
	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed")
	}
}