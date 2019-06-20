package v1

import (
	"github.com/gin-gonic/gin"
	"iads/server/middlewares"
	"iads/server/routers/api/v1/hardware"
	"iads/server/routers/api/v1/user"
	"net/http"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		rbac := middlewares.RbacInitReturnEnforcer()
		v1.Use(middlewares.NewAuthorizer(rbac))
		// 用户模块路由
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"app":     "iads",
				"version": "v1.0",
			})
		})
		v1.GET("/cpuinfo", hardware.CpuInfo)
		user.RegisterUserRouter(v1.Group("/user"))
		user.RegisterRoleRouter(v1.Group("/role"))
	}
}
