package main

import (
	_ "golearning-goframe/internal/packed"

	_ "golearning-goframe/internal/logic"

	_ "golearning-goframe/internal/boot"

	"github.com/gogf/gf/v2/os/gctx"

	"golearning-goframe/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
