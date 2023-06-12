package main

import (
	"github.com/v1lezz/tg-bot-ozon-route/config"
	"github.com/v1lezz/tg-bot-ozon-route/internal/clients/tg"
	"github.com/v1lezz/tg-bot-ozon-route/internal/model/messages"
	"log"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}
	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed", err)
	}
	msgModel := messages.New(tgClient)
	tgClient.ListenUpdates(msgModel)
}
