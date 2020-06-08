package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pkgo/kan/vo"
)

type Policy func(*gin.Context) bool

func Auth(c *gin.Context, policies ...Policy) bool {
	for _, value := range policies {
		b := value(c)
		if b {
			return true
		}
	}
	c.JSON(401, vo.Error("未授权"))
	return false
}


