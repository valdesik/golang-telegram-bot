package telegram

import (
	"telegram-bot-golang/clients/telegram"
	"telegram-bot-golang/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}
