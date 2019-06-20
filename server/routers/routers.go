package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	v1 "iads/server/routers/api/v1"
	"iads/server/routers/api/v1/user"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.POST("/login", user.Auth.LoginHandler)
	api := router.Group("/api")
	{
		v1.RegisterRouter(api)
	}
	return router
}
