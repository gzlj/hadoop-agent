package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/hadoop-agent/pkg/global"
	"github.com/gzlj/hadoop-agent/pkg/infra"
	"github.com/gzlj/hadoop-agent/pkg/module"
)

func HandleGetComponentStatuses(c *gin.Context) {
	var (
		response global.Response
		status module.ClusteredComponentStatuses
		err error
	)
	status, err = infra.GetComponentStatus()
	if err != nil {
		response = global.BuildResponse(500, err.Error(), status)
	} else {
		response = global.BuildResponse(200, "OK", status)
	}
	c.JSON(200,response)
}
