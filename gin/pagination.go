package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var PageSize = 10

func GetGinPage(c *gin.Context) (int, int) {
	offPage := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		offPage = (page - 1) * PageSize
	}
	if page == 0 {
		page = 1
	}
	return offPage, page
}
