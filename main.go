package main

import (
	"flag"
	"fmt"
	"telegram-bot-golang/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	token := mustToken()
	fmt.Println(token)
	//token = flag.Get(token)
	tgClient = telegram.New(tgBotHost, token)
	// fetcher = fetcher.New()
	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token", "", "telegram bot token")
	flag.Parse()
	if *token == "" {
		panic("token is required")
	}
	return *token
}
