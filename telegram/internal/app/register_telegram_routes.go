/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// handlers for telegram commands.
const (
	handleCommandStart = "start"
)

// registerTelegramRouter register telegram routes.
func (a App) registerTelegramRouter(message *tgbotapi.Message) error {
	switch message.Command() {
	case handleCommandStart:
		if err := a.telegramHandler.Start(message); err != nil {
			a.logger.Errorf("🔴 failed to handle command %s: %v", handleCommandStart, err)
		}
	default:
		if err := a.telegramHandler.Default(message); err != nil {
			a.logger.Errorf("🔴 failed to handle command default: %v", err)
		}
	}

	return nil
}
