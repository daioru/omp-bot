package box

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyBoxCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	ok, err := c.boxService.Remove(idx)
	if err != nil {
		log.Printf("fail to delete product with idx %d: %v", idx, err)
		return
	}

	var msg tgbotapi.MessageConfig
	if ok {
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("product with idx %d deleted successfully", idx),
		)
	} else {
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("product with idx %d not deleted", idx),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BoxCommander.Delete: error sending reply message to chat - %v", err)
	}
}
