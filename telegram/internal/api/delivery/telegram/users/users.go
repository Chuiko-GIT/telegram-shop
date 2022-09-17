/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package users

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// handler - define telegram handler.
type handler struct {
	bot *tgbotapi.BotAPI
}

// NewHandler - constructor telegram handler.
func NewHandler(bot *tgbotapi.BotAPI) *handler {
	return &handler{bot: bot}
}

// RegistrationUser - handle registration user command.
func (h handler) RegistrationUser(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "\nЭтот бот 🤖 пока не готов к работе 😿")
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}
