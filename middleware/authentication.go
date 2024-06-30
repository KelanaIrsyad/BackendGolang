package middleware

import (
	"belajar/golang/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unanthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
