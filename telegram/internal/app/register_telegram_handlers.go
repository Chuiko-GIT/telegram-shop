/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import "telegram/internal/api/delivery/telegram"

// registerTelegramHandlers - register telegram handlers.
func (a *App) registerTelegramHandlers() {
	a.telegramHandler = telegram.NewHandler(a.telegramBot)
}
