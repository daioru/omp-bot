package box

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *DummyBoxCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	// Добавить логику переключения страниц
	products, err := c.boxService.List(0, 5)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: 0,
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
