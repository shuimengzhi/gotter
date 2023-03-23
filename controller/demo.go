package controller

import (
	"github.com/gin-gonic/gin"
	"gotter/enum"
	"gotter/request_struct"
	"gotter/service"
	"net/http"
)

func DemoController(c *gin.Context) {
	var demoRequest request_struct.DemoRequest
	err := c.ShouldBindJSON(&demoRequest)
	//utility.VarDumpStop(err.Error())
	// 输入参数有问题报错
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": enum.CodeParamError, "message": err.Error()})
		return
	}
	result := service.DemoService(demoRequest)
	c.JSON(http.StatusOK, gin.H{"code": result.Code, "message": result.Data})
	return
}
