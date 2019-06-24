package v1

import (
	"github.com/gin-gonic/gin"
	"iads/server/routers/api/v1/user"
	"net/http"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		// 用户模块路由
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"app":     "iads",
				"version": "v1.0",
			})
		})
		user.RegisterUserRouter(v1.Group("/user"))
		user.RegisterRoleRouter(v1.Group("/role"))
	}
}
