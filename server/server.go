package server

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	setApi(r.Group("/api"))
}
