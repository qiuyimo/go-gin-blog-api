package app

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuyuhome/go-gin-blog-api/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

// todo.
