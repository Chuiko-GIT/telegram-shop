/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

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
	msg := tgbotapi.NewMessage(message.Chat.ID, "hello")
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

// Default - handle default command.
func (h handler) Default(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "sorry, I don't understand you")
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
