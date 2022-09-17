/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gofiber/fiber/v2"
)

type (
	// StatusHTTPHandler - describes an interface for work with service status over HTTP.
	StatusHTTPHandler interface {
		// CheckStatus - return service status.
		CheckStatus(ctx *fiber.Ctx) error
	}

	// InfoTelegramHandler - describes an interface for work with telegram bot.
	InfoTelegramHandler interface {
		Start(message *tgbotapi.Message) error
		Default(message *tgbotapi.Message) error
		Info(message *tgbotapi.Message) error
	}

	// UsersTelegramHandler - describes an interface for work with telegram bot.
	UsersTelegramHandler interface {
		RegistrationUser(message *tgbotapi.Message) error
	}

	BinanceTelegramHandler interface {
		CryptoCourseToUSDT(message *tgbotapi.Message) error
	}
)
