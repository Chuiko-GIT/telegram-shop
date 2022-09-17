/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"telegram/internal/app"
	"telegram/internal/config"
)

// Block information for the application.
var (
	appName       = "telegram-service"
	version       string
	commit        string
	tag           string
	date          string
	fortuneCookie string
)

// main - define main function for starting service.
func main() {
	cfgPath := flag.String("c", config.DefaultPath, "configuration file path")
	flag.Parse()

	app.NewApp(app.Meta{
		Info: app.Info{
			AppName:      appName,
			Tag:          tag,
			BuildCommit:  commit,
			BuildVersion: version,
			BuildDate:    date,
			BuildCookie:  fortuneCookie,
		},
		ConfigPath: *cfgPath,
	}).Run(registerGracefulHandle())
}

// registerGracefulHandle - define function for register graceful shutdown.
func registerGracefulHandle() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
