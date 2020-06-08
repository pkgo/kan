package token

import (
	"github.com/gin-gonic/gin"
	"github.com/pkgo/kan/vo"
)

type PermissionPolicy func(*gin.Context) bool

func Permission(c *gin.Context, policies ...PermissionPolicy) bool {
	for _, value := range policies {
		b := value(c)
		if b {
			return true
		}
	}
	c.JSON(401, vo.Error("未授权"))
	return false
}