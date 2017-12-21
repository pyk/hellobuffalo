package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pyk/hellobuffalo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
