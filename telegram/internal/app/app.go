/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

import (
	"context"
	"telegram/internal/api/delivery"
	"telegram/internal/config"
	responder "telegram/pkg/http/respoder"
	"telegram/pkg/log"
)

type (
	// Meta - define application meta information.
	Meta struct {
		Info       Info
		ConfigPath string
	}

	// Info defines metadata of application.
	Info struct {
		AppName      string
		Tag          string
		BuildCommit  string
		BuildVersion string
		BuildDate    string
		BuildCookie  string
	}

	// App - define application model.
	App struct {
		// Meta information about application.
		Meta

		// Tech dependencies.
		config *config.Config

		responder responder.Responder
		logger    log.Logger

		// Delivery dependencies.
		statusHTTPHandler delivery.StatusHTTPHandler

		// Services dependencies.

		// Repository dependencies.
	}

	// worker - define worker func.
	worker func(ctx context.Context, a *App)
)

// NewApp - constructor App.
func NewApp(meta Meta) *App {
	return &App{Meta: meta}
}

func (a *App) Run(ctx context.Context) {
	// Initialize configuration
	a.initConfig()

	// Register Dependencies
	a.initLogger()

	// Domain registration.

	// Register Handlers
	a.registerHTTPHandlers()

	// Run Workers
	a.runWorkers(ctx)
}
