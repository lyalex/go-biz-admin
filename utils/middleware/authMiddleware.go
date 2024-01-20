package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyalex/go-biz-admin/utils"
)

func IsAuthenticated(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")
	if _, err := utils.ParseJWt(cookie); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated"},
		)
	}
	return
}
