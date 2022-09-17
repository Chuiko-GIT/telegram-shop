/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram/internal/api/delivery/telegram/binance"
)

// handlers for telegram commands.
const (
	handleCommandStart        = "start"
	handleCommandDefault      = "default"
	handleCommandRegistration = "registration"
	handleCommandInfo         = "info"
)

// registerTelegramRouter register telegram routes.
func (a App) registerTelegramRouter(message *tgbotapi.Message) error {
	// Handle telegram commands.
	switch message.Command() {
	case handleCommandStart:
		if err := a.infoTelegramHandler.Start(message); err != nil {
			a.logger.Errorf("🔴 failed to handle command %s: %v", handleCommandStart, err)
		}
		return nil
	case handleCommandInfo:
		if err := a.infoTelegramHandler.Info(message); err != nil {
			return fmt.Errorf("🔴 failed to handle command %s: %v", handleCommandRegistration, err)
		}
		return nil
	case handleCommandDefault:
		if err := a.infoTelegramHandler.Default(message); err != nil {
			return fmt.Errorf("🔴 failed to handle default command: %v", err)
		}
		return nil
	}

	// Handle telegram messages.
	if binance.SearchCrypto(message.Text).Validate() {
		if err := a.binanceTelegramHandler.CryptoCourseToUSDT(message); err != nil {
			return fmt.Errorf("🔴 failed to handle default binance messages: %v", err)
		}
	}

	return nil
}
