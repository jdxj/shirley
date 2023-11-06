package main

import (
	_ "shirley/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shirley/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
