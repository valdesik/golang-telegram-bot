package telegram

import (
	"log"
	"net/url"
	"strings"
	"telegram-bot-golang/lib/errors"
	"telegram-bot-golang/storage"
)

func (p *Processor) doCmd(text string, ChatId int, username string) {
	text = strings.TrimSpace(text)
	log.Printf("got new commnd '%s' from user %s", text, username)

	if isAddCmd(text) {

	}

}
func (p *Processor) savePage(chatId int, pageURL string, username string) (err error) {
	defer func() { err = errors.WrapIfErr("failed to save page", err) }()
	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}
	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}
	if isExists {
		return nil
	}
	return p.tg.SendMessage(chatId, "saved")

}

func isAddCmd(text string) bool {
	return isUrl(text)
}

func isUrl(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
