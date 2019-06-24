package conf

import (
	"github.com/gin-gonic/gin"
	"time"
)

var (
	AppMaxProc               = 2
	GinMode                  = gin.ReleaseMode
	HttpServerReadTimeout    = 10 * time.Second
	HttpServerWriteTimeout   = 10 * time.Second
	HttpServerMaxHeaderBytes = 1 << 2

	EnableCasbin     = true
	CasbinPath       = "conf/casbin_model.conf"
	CasbinPolicyPath = "conf/policy.csv"
)
