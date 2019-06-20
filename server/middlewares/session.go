package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SessionsMiddleware .
func SessionsMiddleware(c *gin.Context) gin.HandlerFunc {
	store := cookie.NewStore([]byte("secret"))
	return sessions.Sessions("mysession", store)
}
