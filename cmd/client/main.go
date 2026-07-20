package main

import (
	"fmt"
	"log"

	"github.com/Vlad-Pisarevskiy/faraway/config"
	client2 "github.com/Vlad-Pisarevskiy/faraway/internal/client"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := client2.NewClient()
	log.Println("Starting client...")
	log.Fatal(client.Run(fmt.Sprintf(":%s", cfg.Port())))
}
