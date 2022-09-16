/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package app

func (a *App) initWorkers() []worker {
	workers := []worker{
		serveHTTP,
	}

	return workers
}
