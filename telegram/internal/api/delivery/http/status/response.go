/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package status

type (
	// response – describes an interface for work with service status over HTTP.
	response struct {
		Name          string `json:"name"`
		Version       string `json:"version"`
		Tag           string `json:"tag"`
		Commit        string `json:"commit"`
		Date          string `json:"date"`
		FortuneCookie string `json:"fortune_cookie"`
	}
)
