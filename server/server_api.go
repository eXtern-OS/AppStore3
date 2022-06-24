package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Message interface{} `json:"msg"`
}

func handleInfo(c *gin.Context) {
	c.String(http.StatusOK, "You are trying to access eXtern OS Store API, see details here: https://github.com/eXtern-OS/AppStore3")
}

func setApi(r *gin.RouterGroup) {
	r.GET("/", handleInfo)
	r.POST("/search", handleSearch)
}
