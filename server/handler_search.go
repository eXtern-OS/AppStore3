package server

import (
	"externos.io/AppStore3/query"
	"externos.io/AppStore3/search"
	"github.com/gin-gonic/gin"
	"net/http"
)


// handleSearch consumes search query and returns either error code or results
func handleSearch(c *gin.Context) {
	var q query.Query
	if err := c.BindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{Message: "Failed to parse JSON: bad request"})
	} else {
		if q.Results == 0 {
			q.Results = 100
		}
		c.JSON(http.StatusOK, ApiResponse{search.Search(q)})
	}
}
