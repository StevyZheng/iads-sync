package ginadmin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"iads/server/config"
	"iads/server/middleware"
	"iads/server/pkg/logger"
	"iads/server/routers/api"
)

// InitWeb 初始化web引擎
func InitWeb(ctx context.Context, obj *Object) *gin.Engine {
	cfg := config.GetGlobalConfig()
	gin.SetMode(cfg.RunMode)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	apiPrefixes := []string{"/api/"}

	// 跟踪ID
	app.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))

	// 访问日志
	app.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))

	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())

	// 跨域请求
	if cfg.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	// 注册/api路由
	api.RegisterRouter(app, obj.Bll, obj.Auth, obj.Enforcer)

	// swagger文档
	if dir := cfg.Swagger; dir != "" {
		app.Static("/swagger", dir)
	}

	// 静态站点
	if dir := cfg.WWW; dir != "" {
		app.Use(middleware.WWWMiddleware(dir))
	}

	return app
}

// InitHTTPServer 初始化http服务
func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	httpCfg := config.GetGlobalConfig().HTTP
	cfg := config.GetGlobalConfig()
	addr := fmt.Sprintf("%s:%d", httpCfg.Host, httpCfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  cfg.HttpServerReadTimeout,
		WriteTimeout: cfg.HttpServerWriteTimeout,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.StartSpan(ctx, "HTTP服务初始化", "ginadmin.InitHTTPServer").Printf("HTTP服务开始启动，地址监听在：[%s]", addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.StartSpan(ctx, "监听HTTP服务", "ginadmin.InitHTTPServer").Errorf(err.Error())
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(httpCfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.StartSpan(ctx, "关闭HTTP服务", "ginadmin.InitHTTPServer").Errorf(err.Error())
		}
	}
}
