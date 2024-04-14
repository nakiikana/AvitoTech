package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
	cache "tools/internals/cache/middleware"
	"tools/internals/cfg"
	"tools/internals/handler"
	"tools/internals/repository"
	"tools/internals/server"
	"tools/internals/service"
)

func main() {
	config, err := cfg.LoadAndStore("./config")

	if err != nil {
		log.Fatalf("Couldn't parse the config file: %v", err)
	}

	db, err := repository.NewPostgresDB(config)
	if err != nil {
		log.Fatalf("Couldn't connect tp DB: %v", err)
	}
	defer db.Close()

	rep := repository.NewRepository(db)
	cache := cache.NewCache(config)
	srvc := service.NewService(rep, cache)
	hdl := handler.NewHandler(srvc)

	srv := new(server.Server)
	go func() {
		if err = srv.Run(hdl.InitRoutes()); err != nil {
			log.Fatalf("Couldn't start the server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	log.Printf("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)

	log.Printf("An error occured when shutting down the server: %v", err)
}
