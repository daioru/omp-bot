package box

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyBoxCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	editBox, err := c.boxModel.NewBoxFromArgs(args)
	if err != nil {
		log.Printf("fail to create box with args %v: %v", args, err)
		return
	}

	// Тоже создаём объет по аргументам команды
	err = c.boxService.Update(editBox.ID, *editBox)
	if err != nil {
		log.Printf("fail to edit box with idx %d: %v", editBox.ID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(
			"Box with id %d edited successfully",
			editBox.ID,
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BoxCommander.Edit: error sending reply message to chat - %v", err)
	}
}
