package onstart

import (
	"fmt"
	"github.com/gotoeasy/glang/cmn"
	"golangchain/pkg/settings"
	"golangchain/router"
	"net/http"
)

func Run() {
	cmn.Info("Http Server Start")
	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        r,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
