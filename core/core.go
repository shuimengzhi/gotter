package core

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gotter/model"
	"log"
	"os"
)

func LoadCore() {
	//读取.env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//gin模式
	switch os.Getenv("GIN_MODE") {
	case "production":
		//开启生产模式,忽视报错
		gin.SetMode(gin.ReleaseMode)
	case "develop":
		//开发模式,显示报错
		gin.SetMode(gin.DebugMode)
	//case "test":
	//	//测试模式,console打印结果
	//	gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	//拼接数据库连接信息
	dsn := model.GetDsn()
	//创建数据库链接
	model.Database(dsn)
}
