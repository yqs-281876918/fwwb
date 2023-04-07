package controller

import (
	"encoding/base64"
	"encoding/json"
	"fwwb/library/live"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateLivceReq struct {
	Base64 string `json:"base64"`
}

func UpdateLive(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusOK, "error")
		return
	}
	var req UpdateLivceReq
	err = json.Unmarshal(data, &req)
	if err != nil {
		c.String(http.StatusOK, "error")
		return
	}
	live.LastestImgBase64 = req.Base64
	c.String(http.StatusOK, "ok")
}

func GetLive(c *gin.Context) {
	imageBuffer, _ := base64.StdEncoding.DecodeString(live.LastestImgBase64)
	_, _ = c.Writer.WriteString(string(imageBuffer))
}
