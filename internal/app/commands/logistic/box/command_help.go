package box

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyBoxCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__logistic__box - print list of commands\n"+
			"/get__logistic__box - get a entity\n"+
			"/list__logistic__box - get a list of your entity\n"+
			"/delete__logistic__box - delete an existing entity\n\n"+
			"/new__logistic__box - create a new entity\n"+
			"/edit__logistic__box - edit a entity\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BoxCommander.Help: error sending reply message to chat - %v", err)
	}
}
