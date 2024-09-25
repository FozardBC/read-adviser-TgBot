package main

import (
	"log"
	tgClient "main/clients/telegram"
	event_consumer "main/consumer/event-consumer"
	"main/events/telegram"
	"main/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	token := "7227580752:AAGqsim4fWUv_rWH9GI4hSV08bnBMPa0WwY"

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, token),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

/* func mustToken() string {
	token := flag.String("token",
		"",
		"token for access to tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specifeted")
	}

	return *token
}
*/
