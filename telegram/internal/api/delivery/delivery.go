/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package delivery

import "github.com/gofiber/fiber/v2"

type (
	// StatusHTTPHandler - describes an interface for work with service status over HTTP.
	StatusHTTPHandler interface {
		// CheckStatus - return service status.
		CheckStatus(ctx *fiber.Ctx) error
	}
)
