package main

import (
	"log"
	"tools/internals/cfg"
	"tools/internals/handler"
	"tools/internals/repository"
	"tools/internals/server"
	"tools/internals/service"
)

func main() {

	config, err := cfg.LoadAndStore("../config")
	if err != nil {
		log.Fatalf("Couldn't parse the config file: %v", err)
	}
	db, err := repository.NewPostgresDB(config)
	if err != nil {
		log.Fatalf("Couldn't connect tp DB: %v", err)
	}
	rep := repository.NewRepository(db)
	srvc := service.NewService(rep)
	hdl := handler.NewHandler(srvc)
	srv := new(server.Server)
	if err = srv.Run(hdl.InitRoutes()); err != nil {
		log.Fatalf("Couldn't start the server: %v", err)
	}
}
