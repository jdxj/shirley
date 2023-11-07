package main

import (
	_ "github.com/jdxj/shirley/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/shirley/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
