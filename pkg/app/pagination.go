package app

import (
	"strconv"

	"BlogService/global"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize <= 0 {
		return global.Config.App.DefaultPageSize
	}
	if pageSize > global.Config.App.MaxPageSize {
		return global.Config.App.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
