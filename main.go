package main

import (
	_ "problemanalysis-backend/internal/packed"

	"problemanalysis-backend/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
