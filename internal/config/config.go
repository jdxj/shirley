package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Shirley shirley
)

type shirley struct {
	Name string `v:"required"`
	Pass string `v:"required"`
	Key  []byte `v:"required"`
	Path string `v:"required"`
}

func init() {
	ctx := gctx.New()
	shirleyCfg := g.Cfg().MustGet(ctx, "shirley").Map()
	err := gconv.Struct(shirleyCfg, &Shirley)
	if err != nil {
		g.Log().Panicf(ctx, "convert shirley config err: %s", err)
	}

	err = g.Validator().Assoc(shirleyCfg).Data(Shirley).Run(ctx)
	if err != nil {
		g.Log().Panicf(ctx, "check shirley config err: %s", err)
	}
}
