package box

import (
	"fmt"
	"log"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BoxCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		log.Println("argument list cannot be empty", args)
		return
	}

	if utf8.RuneCountInString(args) > 100 {
		log.Println("title field cannot be over 100 characters long")
		return
	}

	newProduct, newProductId, err := c.subdomainService.New(args)
	if err != nil {
		log.Printf("failed to create product with args: %v", args)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(
			"Product with title: %s created successfully. It's id: %d",
			newProduct.Title,
			newProductId,
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BoxCommander.New: error sending reply message to chat - %v", err)
	}
}
