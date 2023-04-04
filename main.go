package main

import (
	bootstarp "fwwb/bootstrap"
	"fwwb/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DangerList(c *gin.Context) {
	res, err := dao.SelectDangers()
	if err != nil {
		c.String(http.StatusOK, "error")
		return
	}
	c.JSON(http.StatusOK, res)
}

func main() {
	bootstarp.InitMysql()
	e := gin.Default() //创建一个默认的路由引擎
	e.GET("/danger/list", DangerList)
	e.Run()
}
