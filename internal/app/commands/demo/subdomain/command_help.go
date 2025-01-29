package subdomain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__logistic__package - print list of commands\n"+
			"/get__logistic__package - get a entity\n"+
			"/list__logistic__package - get a list of your entity\n"+
			"/delete__logistic__package - delete an existing entity\n\n"+
			"/new__logistic__package - create a new entity\n"+
			"/edit__logistic__package - edit a entity\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
