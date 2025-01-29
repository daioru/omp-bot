package box

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyBoxCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argsParts := strings.SplitN(args, " ", 4)
	if len(argsParts) < 4 {
		log.Println("Invalid number of arguments", args)
		return
	}

	newBoxID, err := strconv.Atoi(argsParts[0])
	if err != nil {
		log.Println("wrong id", args)
		return
	}

	weight, err := strconv.ParseFloat(argsParts[1], 32)
	if err != nil {
		log.Println("wrong weight", args)
		return
	}

	volume, err := strconv.ParseFloat(argsParts[2], 32)
	if err != nil {
		log.Println("wrong volume", args)
		return
	}

	isFragile, err := strconv.ParseBool(argsParts[3])
	if err != nil {
		log.Println("wrong isFragile", args)
		return
	}

	//Сгенерировать объект Box из агрументов

	newBox := c.boxModel.NewBox(newBoxID, float32(weight), float32(volume), isFragile)

	newProductID, err := c.boxService.Create(*newBox)
	if err != nil {
		log.Printf("failed to create product with args: %v", args)
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
