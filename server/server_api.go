package server

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Message interface{} `json:"msg"`
}

func setApi(r *gin.RouterGroup) {
	r.POST("/search", handleSearch)
}
