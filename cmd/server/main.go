package main

import (
	"fmt"
	"log"

	"github.com/Vlad-Pisarevskiy/faraway/config"
	"github.com/Vlad-Pisarevskiy/faraway/internal/quotes"
	"github.com/Vlad-Pisarevskiy/faraway/internal/server"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	quoter := quotes.NewQuoter()
	srv := server.NewServer(quoter, cfg.Difficulty())

	log.Println("Starting server...")
	log.Fatal(srv.Run(fmt.Sprintf(":%s", cfg.Port())))
}
