/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import "telegram/internal/api/delivery/http/status"

// registerHTTPHandlers - register http handlers.
func (a *App) registerHTTPHandlers() {
	a.statusHTTPHandler = status.NewHandler(
		a.Meta.Info.AppName,
		a.Meta.Info.Tag,
		a.Meta.Info.BuildVersion,
		a.Meta.Info.BuildCommit,
		a.Meta.Info.BuildDate,
		a.Meta.Info.BuildCookie,
	)
}
