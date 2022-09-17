/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (a *App) initTelegramBot() {
	var err error
	a.telegramBot, err = tgbotapi.NewBotAPI(a.config.Delivery.TelegramBot.Token)
	if err != nil {
		a.logger.Fatalf("🔴 failed to start telegram: %v", err)
	}

	a.telegramBot.Debug = true
}
