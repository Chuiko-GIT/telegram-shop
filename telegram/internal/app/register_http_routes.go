/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import "github.com/gofiber/fiber/v2"

// registerHTTPRoutes - register http routes.
func (a *App) registerHTTPRoutes(app *fiber.App) {
	router := app.Group("/v1/telegram")
	router.Get("/status", a.statusHTTPHandler.CheckStatus)
}
