package box

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *DummyBoxCommander) List(inputMessage *tgbotapi.Message, cursor int, offset int) {
	outputMsgText := fmt.Sprintf("Here all the products for page %d: \n\n", cursor/5 + 1)

	// Добавить логику переключения страниц
	products, err := c.boxService.List(cursor, offset)
	if err != nil {
		log.Printf("Failed to fetch boxes with cursor %d, offset %d: %v", cursor, offset, err)
		return
	}

	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: cursor + 5,
		Limit:  5,
	})

	callbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "box",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BoxCommander.List: error sending reply message to chat - %v", err)
	}
}
