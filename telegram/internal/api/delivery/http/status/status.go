/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package status

import (
	"github.com/gofiber/fiber/v2"
)

// handler - define status model.
type handler struct {
	response
}

// NewHandler - constructor status handler.
func NewHandler(appName, tag, buildVersion, buildCommit, buildDate, fortuneCookie string) *handler {
	return &handler{
		response{
			Name:          appName,
			Version:       buildVersion,
			Tag:           tag,
			Commit:        buildCommit,
			Date:          buildDate,
			FortuneCookie: fortuneCookie,
		},
	}
}

// CheckStatus - HTTP GET handler for status endpoint.
func (h handler) CheckStatus(ctx *fiber.Ctx) error {
	return ctx.JSON(h.response)
}
