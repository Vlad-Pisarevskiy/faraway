package main

import (
	"log"

	"github.com/Vlad-Pisarevskiy/faraway/internal/quotes"
	"github.com/Vlad-Pisarevskiy/faraway/internal/server"
)

func main() {

	quoter := quotes.NewQuoter()
	srv := server.NewServer(quoter, 5)

	log.Fatal(srv.Run(":8080"))
}
