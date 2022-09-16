/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import (
	std "log"

	"telegram/internal/config"
)

func (a *App) initConfig() {
	var err error

	// Cannot assign *Config to a.config (type config.Config) in multiple assignment
	if a.config, err = config.NewConfig(a.Meta.Info.AppName, a.Meta.ConfigPath); err != nil {
		std.Fatal(err)
	}
}
