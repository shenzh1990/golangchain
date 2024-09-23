package onstart

import (
	"fmt"
	"golangchain/pkg/settings"

	"github.com/gotoeasy/glang/cmn"
	"os"
)

func init() {
	cmn.Info("Daemon init")
	httpPort := cmn.IntToString(settings.HTTPPort)
	// 端口冲突时退出
	if cmn.IsPortOpening(httpPort) {
		fmt.Printf("port %s conflict, startup failed.\n", httpPort)
		os.Exit(0)
	}
}
