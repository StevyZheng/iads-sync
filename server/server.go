package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iads/server/config"
	"iads/server/routers"
	"net/http"
	"runtime"
)

func ServerStart() {
	cfg := config.GetGlobalConfig()
	runtime.GOMAXPROCS(cfg.AppMaxProc)
	gin.SetMode(cfg.GinMode)
	router := routers.InitRouter()

	ser := &http.Server{
		Addr:           fmt.Sprintf(":%d", 80),
		Handler:        router,
		ReadTimeout:    cfg.HttpServerReadTimeout,
		WriteTimeout:   cfg.HttpServerWriteTimeout,
		MaxHeaderBytes: cfg.HttpServerMaxHeaderBytes,
	}
	_ = ser.ListenAndServe()
}
