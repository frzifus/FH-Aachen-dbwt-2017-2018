package util

import (
	"github.com/gernest/utron/base"
)

func CtxTitle(t string) func(*base.Context) error {
	return func(ctx *base.Context) error {
		ctx.Data["title"] = t
		return nil
	}
}
