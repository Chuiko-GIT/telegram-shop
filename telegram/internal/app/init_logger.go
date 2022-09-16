/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import (
	"telegram/pkg/log"
)

// initLogger initializes logger.
func (a *App) initLogger() {
	a.logger = log.InitLogger(a.config.Logger, map[string]string{
		"service": a.Meta.Info.AppName,
	})
}
