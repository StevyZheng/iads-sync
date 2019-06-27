package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
	"iads/server/config"
	"iads/server/pkg/logger"
	"runtime"

	. "iads/server/init"
)

func ServerStart(ctx context.Context) func() {
	cfg := config.GetGlobalConfig()
	runtime.GOMAXPROCS(cfg.AppMaxProc)
	gin.SetMode(cfg.GinMode)

	loggerCall, err := InitLogger()
	if err != nil {
		panic(err)
	}

	if c := config.GetGlobalConfig().Monitor; c.Enable {
		err = agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: true})
		if err != nil {
			logger.StartSpan(ctx, "开启[agent]服务监听", "ginadmin.Init").Errorf(err.Error())
		}
	}

	InitCaptcha()

	obj, objCall, err := InitObject(ctx)
	if err != nil {
		panic(err)
	}

	err = InitData(ctx, obj)
	if err != nil {
		logger.StartSpan(ctx, "初始化应用数据", "ginadmin.Init").Errorf(err.Error())
	}

	app := InitWeb(ctx, obj)
	httpCall := InitHTTPServer(ctx, app)

	return func() {
		if httpCall != nil {
			httpCall()
		}

		if objCall != nil {
			objCall()
		}

		if loggerCall != nil {
			loggerCall()
		}
	}
}
