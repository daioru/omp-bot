package box

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyBoxCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	newBox, err := c.boxModel.NewBoxFromArgs(args)
	if err != nil {
		log.Printf("fail to create box with args %v: %v", args, err)
		return
	}

	newProductID, err := c.boxService.Create(*newBox)
	if err != nil {
		log.Printf("failed to store box with args %v: %v", args, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(
			"Box with id %d sussessfully created",
			newProductID,
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BoxCommander.New: error sending reply message to chat - %v", err)
	}
}
