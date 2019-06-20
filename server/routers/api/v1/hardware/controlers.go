package hardware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"iads/lib/linux"
	"net/http"
)

func CpuInfo(c *gin.Context) {
	cpu := linux.CpuHwInfo{}
	cpu.GetCpuHwInfo()
	jsons, errs := json.MarshalIndent(cpu, "", "  ")
	if errs != nil {
	}
	c.String(http.StatusOK, string(jsons))
}
