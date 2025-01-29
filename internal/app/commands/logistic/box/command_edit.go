package box

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyBoxCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argsParts := strings.SplitN(args, " ", 4)
	if len(argsParts) < 4 {
		log.Println("Invalid number of arguments", args)
		return
	}

	editBoxID, err := strconv.Atoi(argsParts[0])
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

	editBox := c.boxModel.NewBox(editBoxID, float32(weight), float32(volume), isFragile)

	// Тоже создаём объет по аргументам команды
	err = c.boxService.Update(editBoxID, *editBox)
	if err != nil {
		log.Printf("fail to edit box with idx %d: %v", editBoxID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(
			"Box with id: %d edited successfully",
			editBoxID,
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BoxCommander.Edit: error sending reply message to chat - %v", err)
	}
}
