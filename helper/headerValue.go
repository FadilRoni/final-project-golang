package helper

import "github.com/gin-gonic/gin"

func GetContent(c *gin.Context) string {
	return c.Request.Header.Get("Content-type")
}
