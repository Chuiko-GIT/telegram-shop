/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import (
	"telegram/internal/api/delivery/telegram/binance"
	"telegram/internal/api/delivery/telegram/info"
	"telegram/internal/api/delivery/telegram/users"
)

// registerTelegramHandlers - register telegram handlers.
func (a *App) registerTelegramHandlers() {
	a.infoTelegramHandler = info.NewHandler(a.telegramBot)
	a.usersTelegramHandler = users.NewHandler(a.telegramBot)
	a.binanceTelegramHandler = binance.NewHandler(a.telegramBot, a.config.Delivery.TelegramBot.Binance)
}
