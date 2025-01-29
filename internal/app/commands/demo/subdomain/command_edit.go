package subdomain

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argsParts := strings.SplitN(args, " ", 2)
	if len(argsParts) < 2 {
		log.Println("Invalid number of arguments", args)
		return
	}

	idx, err := strconv.Atoi(argsParts[0])
	if err != nil {
		log.Println("wrong id format", args)
		return
	}

	editedProduct, err := c.subdomainService.Edit(idx, argsParts[1])
	if err != nil {
		log.Printf("fail to edit product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(
			"Product with id: %d edited successfully. New title: %s",
			idx,
			editedProduct.Title,
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Edit: error sending reply message to chat - %v", err)
	}
}
