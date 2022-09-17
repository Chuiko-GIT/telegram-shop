/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package info

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// handler - define telegram handler.
type handler struct {
	bot *tgbotapi.BotAPI
}

// NewHandler - constructor telegram handler.
func NewHandler(bot *tgbotapi.BotAPI) *handler {
	return &handler{bot: bot}
}

// Start - handle start command.
func (h handler) Start(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"Привет, %s 🖐"+
			"\nДобро пожаловать в наш бот! \U0001F9BE🤖🖖", message.From.UserName),
	)

	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

// Info - handle info command.
func (h handler) Info(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"На данный момент, бот умеет работать с курсом криптовалюты к USDT.🤖"+
			"\nДля получения курса введите название криптовалюты [BTC, ETH, XLM]. 🤖"+
			"\nНапример: BTC"),
	)

	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

// Default - handle default command.
func (h handler) Default(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"Извини, я не умею работать с этой командой 😿: \n%s", message.Text),
	)

	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}
