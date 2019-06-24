package middleware

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	. "iads/server/conf"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !EnableCasbin {
			c.Next()
			return
		}
		p := c.Request.URL.Path
		m := c.Request.Method
	}
}
