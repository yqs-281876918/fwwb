package main

import (
	"encoding/json"
	bootstarp "fwwb/bootstrap"
	"fwwb/controller"
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

func DangerAdd(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusOK, "error")
		return
	}
	var body dao.Danger
	err = json.Unmarshal(data, &body)
	if err != nil {
		c.String(http.StatusOK, "error")
		return
	}
	err = dao.InsertDanger(body)
	if err != nil {
		c.String(http.StatusOK, "error")
		return
	}
	c.String(http.StatusOK, "ok")
}

func main() {
	bootstarp.InitMysql()
	e := gin.Default() //创建一个默认的路由引擎
	e.GET("/danger/list", DangerList)
	e.GET("/live/img", controller.GetLive)
	e.POST("/live/update", controller.UpdateLive)
	e.Run()
}
