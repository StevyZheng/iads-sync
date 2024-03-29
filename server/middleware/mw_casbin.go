package middleware

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"iads/server/config"
	"iads/server/ginplus"
	"iads/server/pkg/errors"
	"iads/server/pkg/logger"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware(enforcer *casbin.Enforcer, skipper ...SkipperFunc) gin.HandlerFunc {
	cfg := config.GetGlobalConfig()
	return func(c *gin.Context) {
		if !cfg.EnableCasbin || len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		p := c.Request.URL.Path
		m := c.Request.Method
		if b, err := enforcer.EnforceSafe(ginplus.GetUserID(c), p, m); err != nil {
			logger.StartSpan(ginplus.NewContext(c), "casbin中间件", "CasbinMiddleware").
				Errorf(err.Error())
			ginplus.ResError(c, errors.NewInternalServerError())
			return
		} else if !b {
			ginplus.ResError(c, errors.NewForbiddenError("没有访问权限"))
			return
		}
		c.Next()
	}
}
