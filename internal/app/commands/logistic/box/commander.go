package box

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/box"
)

type BoxCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type DummyBoxCommander struct {
	bot        *tgbotapi.BotAPI
	boxService *box.DummyBoxService
	boxModel   *box.DummyBoxModel
}

func NewDummyBoxCommander(
	bot *tgbotapi.BotAPI,
) *DummyBoxCommander {
	boxService := box.NewDummyBoxService()

	return &DummyBoxCommander{
		bot:        bot,
		boxService: boxService,
	}
}

func (c *DummyBoxCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("BoxCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DummyBoxCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
