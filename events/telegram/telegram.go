package telegram

import "telegram-bot-golang/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
}
