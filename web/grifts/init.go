package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/leo1971/pikachu/web/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
