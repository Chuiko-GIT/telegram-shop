/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func serveTelegram(ctx context.Context, a *App) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := a.telegramBot.GetUpdatesChan(u)
	if err != nil {
		a.logger.Fatalf("🔴 GetUpdatesChan telegram: %v", err)
	}

	for {
		select {
		case update := <-updates:
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			// Handle commands
			if update.Message.IsCommand() {
				a.registerTelegramRouter(update.Message)
				continue
			}

			//// Handle regular messages
			//if err := botApi.handleMessage(update.Message); err != nil {
			//	botApi.handleError(update.Message.Chat.ID, err)
			//}

		case <-ctx.Done():
			a.logger.Info("🔴 telegram bot gracefully stopped")
			return
		}
	}
}
