package middleware

import (
	"final-project/helper"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			helper.Unauthorized(c, "failed to verify token")
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
