package commons

import "github.com/gin-gonic/gin"

func GetOperatorIDFromHeader(c *gin.Context) string {
	return c.GetHeader("x-operator-id")
}
