package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "iads/server/conf"
	"iads/server/routers"
	"net/http"
	"runtime"
)

func ServerStart() {
	runtime.GOMAXPROCS(AppMaxProc)
	gin.SetMode(GinMode)
	router := routers.InitRouter()

	ser := &http.Server{
		Addr:           fmt.Sprintf(":%d", 80),
		Handler:        router,
		ReadTimeout:    HttpServerReadTimeout,
		WriteTimeout:   HttpServerWriteTimeout,
		MaxHeaderBytes: HttpServerMaxHeaderBytes,
	}
	_ = ser.ListenAndServe()
}
